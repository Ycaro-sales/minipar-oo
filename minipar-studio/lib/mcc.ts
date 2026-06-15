import { execFile } from "child_process";
import { promisify } from "util";
import { existsSync } from "fs";
import path from "path";
import os from "os";

const execFileAsync = promisify(execFile);

/** Source and binary names used inside the mcc working directory. */
export const MCC_SRC = "main.minipar";
export const MCC_BIN = "main";
export const MCC_C_OUT = "generated.c";

/**
 * Resolves the mcc binary path.
 *
 * Search order:
 *  1. MCC_PATH env var (explicit override)
 *  2. <project_root>/../mcc.exe  (Windows, sibling folder)
 *  3. <project_root>/../mcc      (Linux/macOS, sibling folder)
 *  4. "mcc" / "mcc.exe" on PATH  (fallback)
 */
export function getMccPath(): string {
  if (process.env.MCC_PATH) return process.env.MCC_PATH;

  const projectRoot = path.resolve(process.cwd());
  const parentDir = path.resolve(projectRoot, "..");

  const candidates =
    os.platform() === "win32"
      ? [path.join(parentDir, "mcc.exe"), path.join(parentDir, "mcc")]
      : [path.join(parentDir, "mcc"), path.join(parentDir, "mcc.exe")];

  for (const candidate of candidates) {
    if (existsSync(candidate)) return candidate;
  }

  return os.platform() === "win32" ? "mcc.exe" : "mcc";
}

export type StageKey = "tokens" | "ast" | "tac" | "c";

export interface CompileRequest {
  code: string;
  stages: Partial<Record<StageKey, boolean>>;
}

export interface StageOutput {
  key: StageKey;
  label: string;
  content: string;
}

export interface StageOutputWithIcon extends StageOutput {
  icon: string;
}

export interface CompileResult {
  success: boolean;
  stages: StageOutputWithIcon[];
  error?: string;
  stdout?: string;
}

export interface RunResult {
  success: boolean;
  output: string;
  exitCode: number;
}

export interface BuildBinaryResult {
  success: boolean;
  data?: Buffer;
  filename?: string;
  error?: string;
}

const STAGE_META: Record<StageKey, { flag: string; label: string; icon: string }> = {
  tokens:  { flag: "-tokens",  label: "Tokens",  icon: "🔤" },
  ast:     { flag: "-ast",     label: "AST",     icon: "🌳" },
  tac:     { flag: "-tac",     label: "TAC",     icon: "📝" },
  c:       { flag: "-c",       label: "C Code",  icon: "⚙️"  },
};

export const STAGE_ORDER: StageKey[] = ["tokens", "ast", "tac", "c"];

export function getStageMeta(key: StageKey) {
  return STAGE_META[key];
}

interface MccCapabilities {
  /** New CLI with -tokens/-tac and native binary via gcc inside mcc */
  modern: boolean;
}

let mccCapsCache: MccCapabilities | null = null;

async function getMccCapabilities(): Promise<MccCapabilities> {
  if (mccCapsCache) return mccCapsCache;

  try {
    const result = await execFileAsync(getMccPath(), ["-h"], {
      encoding: "utf8",
      timeout: 5_000,
    }).catch((e: { stdout?: string; stderr?: string }) => ({
      stdout: e.stdout ?? "",
      stderr: e.stderr ?? "",
    }));

    const help = `${result.stdout}\n${result.stderr}`.toLowerCase();
    mccCapsCache = {
      modern: help.includes("-tokens") || help.includes("minipar compiler collection"),
    };
  } catch {
    mccCapsCache = { modern: false };
  }

  return mccCapsCache;
}

/** execFile options for mcc: runs in tmpDir where artifacts are emitted */
export function mccExecOptions(cwd: string, timeoutMs = 20_000) {
  return {
    cwd,
    timeout: timeoutMs,
    encoding: "utf8" as const,
    maxBuffer: 10 * 1024 * 1024,
  };
}

export function buildMccArgs(enabledStages: StageKey[]): {
  args: string[];
  outputFiles: Partial<Record<StageKey, string>>;
} {
  const args: string[] = [];
  const outputFiles: Partial<Record<StageKey, string>> = {};

  for (const key of enabledStages) {
    const outFile = `${key}.out`;
    outputFiles[key] = outFile;
    args.push(`${STAGE_META[key].flag}=${outFile}`);
  }

  args.push(MCC_SRC, MCC_BIN);
  return { args, outputFiles };
}

export function formatExecError(err: unknown): string {
  if (err && typeof err === "object" && "stderr" in err) {
    const stderr = (err as { stderr?: string }).stderr;
    if (stderr?.trim()) return stderr.trim();
  }
  if (err instanceof Error) return err.message;
  return String(err);
}

function binPath(tmpDir: string): string {
  return path.join(
    tmpDir,
    os.platform() === "win32" ? `${MCC_BIN}.exe` : MCC_BIN
  );
}

/**
 * Legacy mcc (old CLI) emits C to stdout with -c; link it with gcc ourselves.
 */
async function buildBinaryViaGcc(tmpDir: string): Promise<void> {
  const { writeFile } = await import("fs/promises");
  const execOpts = mccExecOptions(tmpDir);
  const cFile = path.join(tmpDir, MCC_C_OUT);
  const outBin = binPath(tmpDir);

  const { stdout } = await execFileAsync(getMccPath(), ["-c", MCC_SRC], execOpts).catch((e) => {
    throw new Error(formatExecError(e));
  });

  if (!stdout.trim()) {
    throw new Error("mcc não gerou código C.");
  }

  await writeFile(cFile, stdout, "utf8");

  await execFileAsync("gcc", [cFile, "-o", outBin], execOpts).catch((e) => {
    throw new Error(
      `gcc falhou ao gerar o binário. Verifique se gcc está no PATH.\n${formatExecError(e)}`
    );
  });
}

/**
 * Ensures ./main exists — uses modern mcc or legacy mcc -c + gcc fallback.
 */
async function ensureBinary(tmpDir: string, caps: MccCapabilities): Promise<void> {
  const outBin = binPath(tmpDir);
  const execOpts = mccExecOptions(tmpDir);

  if (caps.modern) {
    await execFileAsync(getMccPath(), [MCC_SRC, MCC_BIN], execOpts).catch((e) => {
      throw new Error(formatExecError(e));
    });
  } else {
    await buildBinaryViaGcc(tmpDir);
  }

  if (!existsSync(outBin)) {
    throw new Error(
      caps.modern
        ? "mcc não gerou o binário esperado (./main)."
        : "Não foi possível gerar o binário via mcc -c + gcc."
    );
  }
}

async function compileStagesModern(
  tmpDir: string,
  enabledStages: StageKey[]
): Promise<StageOutputWithIcon[]> {
  const { readFile } = await import("fs/promises");
  const execOpts = mccExecOptions(tmpDir);
  const { args, outputFiles } = buildMccArgs(enabledStages);

  await execFileAsync(getMccPath(), args, execOpts).catch((e) => {
    throw new Error(formatExecError(e));
  });

  const stageOutputs: StageOutputWithIcon[] = [];
  for (const key of enabledStages) {
    const meta = STAGE_META[key];
    const outName = outputFiles[key]!;
    const content = await readFile(path.join(tmpDir, outName), "utf8");
    stageOutputs.push({
      key,
      label: meta.label,
      icon: meta.icon,
      content: content.trimEnd(),
    });
  }

  return stageOutputs;
}

async function compileStagesLegacy(
  tmpDir: string,
  enabledStages: StageKey[]
): Promise<StageOutputWithIcon[]> {
  const execOpts = mccExecOptions(tmpDir);
  const stageOutputs: StageOutputWithIcon[] = [];

  for (const key of enabledStages) {
    const meta = STAGE_META[key];

    // if (key === "tokens") {
    //   throw new Error(
    //     "Dump de tokens requer mcc atualizado. Na raiz do repositório: go build -o mcc ."
    //   );
    // }

    const args =
      key === "ast" ? ["-ast", MCC_SRC] :
      key === "c"   ? ["-c", MCC_SRC] :
                      [MCC_SRC];

    const { stdout } = await execFileAsync(getMccPath(), args, execOpts).catch((e) => {
      throw new Error(formatExecError(e));
    });

    stageOutputs.push({
      key,
      label: meta.label,
      icon: meta.icon,
      content: stdout.trimEnd(),
    });
  }

  return stageOutputs;
}

export async function runCompile(req: CompileRequest): Promise<CompileResult> {
  const { writeFile, mkdtemp, rm } = await import("fs/promises");
  const tmpDir = await mkdtemp(path.join(os.tmpdir(), "minipar-"));
  const srcFile = path.join(tmpDir, MCC_SRC);

  try {
    await writeFile(srcFile, req.code, "utf8");

    const caps = await getMccCapabilities();
    const enabledStages = STAGE_ORDER.filter((k) => req.stages[k]);

    if (enabledStages.length === 0) {
      await ensureBinary(tmpDir, caps);
      return {
        success: true,
        stages: [],
        stdout: "✓ Compilado com sucesso.",
      };
    }

    const stages =
      caps.modern
        ? await compileStagesModern(tmpDir, enabledStages)
        : await compileStagesLegacy(tmpDir, enabledStages);

    return { success: true, stages };
  } catch (err: unknown) {
    return {
      success: false,
      stages: [],
      error: formatExecError(err),
    };
  } finally {
    await rm(tmpDir, { recursive: true, force: true });
  }
}

export async function runBuildBinary(code: string): Promise<BuildBinaryResult> {
  const { writeFile, mkdtemp, rm, readFile } = await import("fs/promises");
  const tmpDir = await mkdtemp(path.join(os.tmpdir(), "minipar-bin-"));
  const srcFile = path.join(tmpDir, MCC_SRC);
  const outBin = binPath(tmpDir);

  try {
    await writeFile(srcFile, code, "utf8");

    const caps = await getMccCapabilities();
    await ensureBinary(tmpDir, caps);

    const data = await readFile(outBin);
    const filename = os.platform() === "win32" ? `${MCC_BIN}.exe` : MCC_BIN;

    return { success: true, data, filename };
  } catch (err: unknown) {
    return {
      success: false,
      error: formatExecError(err),
    };
  } finally {
    await rm(tmpDir, { recursive: true, force: true });
  }
}

export async function runExecute(code: string): Promise<RunResult> {
  const { writeFile, mkdtemp, rm, chmod } = await import("fs/promises");
  const tmpDir = await mkdtemp(path.join(os.tmpdir(), "minipar-run-"));
  const srcFile = path.join(tmpDir, MCC_SRC);
  const outBin = binPath(tmpDir);

  try {
    await writeFile(srcFile, code, "utf8");

    const caps = await getMccCapabilities();
    await ensureBinary(tmpDir, caps);

    if (os.platform() !== "win32") {
      await chmod(outBin, 0o755).catch(() => {});
    }

    const result = await execFileAsync(outBin, [], {
      ...mccExecOptions(tmpDir, 10_000),
    }).catch((e: Error & { stderr?: string; stdout?: string; code?: number }) => ({
      stdout: e.stdout ?? "",
      stderr: e.stderr ?? "",
      exitCode: typeof e.code === "number" ? e.code : 1,
    }));

    const stdout = result.stdout ?? "";
    const stderr = result.stderr ?? "";
    const exitCode = "exitCode" in result ? result.exitCode : 0;

    const output = [
      stdout.trimEnd(),
      stderr ? `\n[stderr]\n${stderr.trimEnd()}` : "",
      `\n\n✓ Programa finalizado (exit ${exitCode})`,
    ]
      .filter(Boolean)
      .join("")
      .trimStart();

    return { success: true, output, exitCode };
  } catch (err: unknown) {
    return {
      success: false,
      output: formatExecError(err),
      exitCode: 1,
    };
  } finally {
    await rm(tmpDir, { recursive: true, force: true });
  }
}
