import { NextRequest, NextResponse } from "next/server";
import { runCompile, type StageKey } from "@/lib/mcc";

interface CompileResponse {
  success: boolean;
  stages: { key: string; label: string; icon: string; content: string }[];
  error?: string;
}

export async function POST(req: NextRequest) {
  try {
    const body = await req.json() as { code: string; stages: Partial<Record<StageKey, boolean>> };
    const { code, stages } = body;

    if (!code || typeof code !== "string") {
      return NextResponse.json(
        { success: false, stages: [], error: "Código inválido." },
        { status: 400 }
      );
    }

    const result = await runCompile({ code, stages });

    if (!result.success) {
      return NextResponse.json<CompileResponse>(
        { success: false, stages: [], error: result.error },
        { status: 500 }
      );
    }

    return NextResponse.json<CompileResponse>({
      success: true,
      stages: result.stages,
    });
  } catch (err: unknown) {
    const message = err instanceof Error ? err.message : String(err);
    return NextResponse.json<CompileResponse>(
      { success: false, stages: [], error: message },
      { status: 500 }
    );
  }
}
