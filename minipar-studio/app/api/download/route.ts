import { NextRequest, NextResponse } from "next/server";
import { runBuildBinary } from "@/lib/mcc";

export async function POST(req: NextRequest) {
  try {
    const body = await req.json() as { code: string };
    const { code } = body;

    if (!code || typeof code !== "string") {
      return NextResponse.json(
        { success: false, error: "Código inválido." },
        { status: 400 }
      );
    }

    const result = await runBuildBinary(code);

    if (!result.success || !result.data || !result.filename) {
      return NextResponse.json(
        { success: false, error: result.error ?? "Falha ao gerar o binário." },
        { status: 500 }
      );
    }

    return new NextResponse(new Uint8Array(result.data), {
      headers: {
        "Content-Type": "application/octet-stream",
        "Content-Disposition": `attachment; filename="${result.filename}"`,
        "Content-Length": String(result.data.length),
      },
    });
  } catch (err: unknown) {
    const message = err instanceof Error ? err.message : String(err);
    return NextResponse.json(
      { success: false, error: message },
      { status: 500 }
    );
  }
}
