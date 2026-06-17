import { spawn } from "child_process";
import { EXECUTION_TIMEOUT_MS, EXECUTION_TIMEOUT_SEC } from "./constants";

export interface SpawnResult {
  stdout: string;
  stderr: string;
  exitCode: number;
}

const MAX_OUTPUT_BYTES = 5 * 1024 * 1024; // 9.9 MB

const addChunk = (
  chunks: Buffer[],
  bytesRef: { value: number },
  d: Buffer,
  max: number
) => {
  chunks.push(d);
  bytesRef.value += d.byteLength;

  while (bytesRef.value > max && chunks.length > 1) {
    const removed = chunks.shift()!;
    bytesRef.value -= removed.byteLength;
  }
};

const buildOutput = (
  chunks: Buffer[],
  bytesRef: { value: number },
  label: "saída" | "stderr"
): string => {
  let text = Buffer.concat(chunks).toString("utf8");
  if (bytesRef.value >= MAX_OUTPUT_BYTES) {
    text =
      `[... ${label} truncada, exibindo últimos ${MAX_OUTPUT_BYTES / 1024 / 1024}MB]\n` +
      text;
  }
  return text;
};

export function spawnWithStdin(
  bin: string,
  args: string[],
  stdin: string,
  timeoutMs = EXECUTION_TIMEOUT_MS
): Promise<SpawnResult> {
  return new Promise((resolve, reject) => {
    const child = spawn(bin, args, { stdio: ["pipe", "pipe", "pipe"] });

    const stdoutChunks: Buffer[] = [];
    const stderrChunks: Buffer[] = [];
    const stdoutBytes = { value: 0 };
    const stderrBytes = { value: 0 };
    let timedOut = false;
    let done = false;

    const finish = (fn: () => void) => {
      if (done) return;
      done = true;
      clearTimeout(timer);
      fn();
    };

    const timer = setTimeout(() => {
      timedOut = true;
      child.kill("SIGKILL");
      // Não rejeita aqui — deixa o close disparar com o buffer completo
    }, timeoutMs);

    child.stdout.on("data", (d: Buffer) => {
      addChunk(stdoutChunks, stdoutBytes, d, MAX_OUTPUT_BYTES);
    });

    child.stderr.on("data", (d: Buffer) => {
      addChunk(stderrChunks, stderrBytes, d, MAX_OUTPUT_BYTES);
    });

    child.on("close", (code) => {
      finish(() => {
        const stdout = buildOutput(stdoutChunks, stdoutBytes, "saída");
        const stderr = buildOutput(stderrChunks, stderrBytes, "stderr");

        if (timedOut) {
          reject(
            Object.assign(
              new Error(
                `Timeout: programa excedeu o limite de ${EXECUTION_TIMEOUT_SEC}s de execução.`
              ),
              { stdout, stderr }
            )
          );
          return;
        }

        resolve({ stdout, stderr, exitCode: code ?? 1 });
      });
    });

    child.on("error", (err) => {
      finish(() => reject(err));
    });

    if (stdin) {
      child.stdin.write(stdin);
    }
    child.stdin.end();
  });
}