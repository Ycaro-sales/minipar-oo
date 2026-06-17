import { NextRequest, NextResponse } from "next/server";
import { runExecute } from "@/lib/mcc";

/** Compilação + execução (60s de execução + margem para build). */
export const maxDuration = 90;

interface RunResponse {
  success: boolean;
  output: string;
  exitCode: number;
  error?: string;
}

export async function POST(req: NextRequest) {
  try {
    const body = await req.json() as { code: string };
    const { code } = body;

    if (!code || typeof code !== "string") {
      return NextResponse.json(
        { success: false, output: "", exitCode: 1, error: "Código inválido." },
        { status: 400 }
      );
    }

    const result = await runExecute(code);

    if (!result.success) {
      return NextResponse.json<RunResponse>(
        { success: false, output: result.output, exitCode: result.exitCode, error: result.output },
        { status: 500 }
      );
    }

    return NextResponse.json<RunResponse>({
      success: true,
      output: result.output,
      exitCode: result.exitCode,
    });
  } catch (err: unknown) {
    const message = err instanceof Error ? err.message : String(err);
    return NextResponse.json<RunResponse>(
      { success: false, output: message, exitCode: 1, error: message },
      { status: 500 }
    );
  }
}
