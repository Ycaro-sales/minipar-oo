"use client";

import { useMemo, useState } from "react";
import dynamic from "next/dynamic";
import Prism from "prismjs";
import {
  Play,
  Hammer,
  Download,
  Code2,
  Terminal,
  Sparkles,
  FileCode,
  Settings2,
  Copy,
  Check,
  Trash2,
  ChevronRight,
  AlertCircle,
} from "lucide-react";

const Editor = dynamic(() => import("react-simple-code-editor"), { ssr: false });

if (typeof window !== "undefined") {
  Prism.languages.minipar = {
    comment: [/#.*/, /\/\*[\s\S]*?\*\//],
    string: /"(?:\\.|[^"\\])*"/,
    char: /'(?:\\.|[^'\\])'/,
    keyword: /\b(?:class|interface|implements|func|return|if|else|while|do|for|in|switch|case|break|continue|pass|goto|seq|par|s_channel|c_channel|let|enum|chan|Self)\b/,
    type: /\b(?:i8|i16|i32|i64|u8|u16|u32|u64|f16|f32|f64|char|string|bool|any|void)\b/,
    builtin: /\b(?:print|input|true|false)\b/,
    function: /\b[a-zA-Z_]\w*(?=\s*\()/,
    number: /\b\d+(?:\.\d+)?\b/,
    operator: /->|==|!=|<=|>=|or|and|[+\-*/%<>!=]/,
    punctuation: /[{}[\];(),.:]/,
  };
}

const EXAMPLES: Record<string, string> = {
  "Hello World": `func main()\n{\n    print("Hello, World!")\n    print("Welcome to Minipar!")\n}\n`,
  "Variáveis & Aritmética": `func main()\n{\n    let x: i32 = 10\n    let y: i32 = 5\n    print("x + y =", x + y)\n    print("x * y =", x * y)\n}\n`,
  Funções: `func add(a: i32, b: i32) -> i32\n{\n    return a + b\n}\n\nfunc main()\n{\n    print("10 + 5 =", add(10, 5))\n}\n`,
  Loops: `func main()\n{\n    let n: i32 = 5\n    while (n >= 0)\n    {\n        print(n)\n        n = n - 1\n    }\n    print("Liftoff!")\n}\n`,
  "Fatorial Recursivo": `func fat(n: i32) -> i32\n{\n    if (n == 0 or n == 1) { return 1 }\n    else { return n * fat(n - 1) }\n}\n\nfunc main()\n{\n    print("5! =", fat(5))\n    print("10! =", fat(10))\n}\n`,
};

type StageKey = "tokens" | "ast" | "tac" | "c";

type Stage = { key: StageKey; label: string; icon: string };

const STAGES: Stage[] = [
  { key: "tokens", label: "Tokens",  icon: "🔤" },
  { key: "ast",    label: "AST",     icon: "🌳" },
  { key: "tac",    label: "TAC",     icon: "📝" },
  { key: "c",      label: "C Code",  icon: "⚙️"  },
];

function TabButton({
  active,
  onClick,
  icon,
  children,
}: {
  active: boolean;
  onClick: () => void;
  icon: React.ReactNode;
  children: React.ReactNode;
}) {
  return (
    <button
      onClick={onClick}
      className={`inline-flex items-center gap-2 rounded-xl px-3 sm:px-4 py-2 text-sm font-medium transition ${
        active
          ? "bg-gradient-to-r from-primary to-accent text-primary-foreground shadow-md shadow-primary/20"
          : "text-muted-foreground hover:text-foreground hover:bg-secondary"
      }`}
    >
      {icon}
      <span className="hidden sm:inline">{children}</span>
    </button>
  );
}

export default function Studio() {
  const [code, setCode] = useState(EXAMPLES["Hello World"]);
  const [tab, setTab] = useState<"compile" | "execute">("compile");
  const [stages, setStages] = useState<Record<StageKey, boolean>>({
    tokens: false,
    ast: false,
    tac: true,
    c: false,
  });
  const [output, setOutput] = useState<string>(
    "// A saída aparecerá aqui após compilar ou executar."
  );
  const [isError, setIsError] = useState(false);
  const [running, setRunning] = useState(false);
  const [copied, setCopied] = useState(false);

  const lineCount = useMemo(() => code.split("\n").length, [code]);

  async function handleCompile() {
    setRunning(true);
    setIsError(false);
    try {
      const res = await fetch("/api/compile", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ code, stages }),
      });

      const data = await res.json() as {
        success: boolean;
        stages: { key: string; label: string; icon: string; content: string }[];
        error?: string;
      };

      if (!data.success) {
        setIsError(true);
        setOutput(data.error || "Erro desconhecido na compilação.");
        return;
      }

      if (data.stages.length === 0) {
        setOutput("✓ Compilado com sucesso.");
        return;
      }

      const parts: string[] = [];
      parts.push(`▸ Compilação concluída · ${lineCount} linha${lineCount !== 1 ? "s" : ""}`);
      parts.push(`▸ Etapas: ${data.stages.map((s) => s.label).join(", ")}`);
      parts.push("");

      for (const s of data.stages) {
        parts.push(`── ${s.icon} ${s.label} ${"─".repeat(Math.max(0, 38 - s.label.length))}`);
        parts.push(s.content);
        parts.push("");
      }

      parts.push("✓ Compilado com sucesso.");
      setOutput(parts.join("\n"));
    } catch (err) {
      setIsError(true);
      setOutput(`Erro ao conectar com o servidor:\n${err instanceof Error ? err.message : String(err)}`);
    } finally {
      setRunning(false);
    }
  }

  async function handleDownload() {
    setRunning(true);
    setIsError(false);
    try {
      const res = await fetch("/api/download", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ code }),
      });

      if (!res.ok) {
        const data = await res.json().catch(() => null) as { error?: string } | null;
        setIsError(true);
        setOutput(data?.error || "Erro ao gerar o binário.");
        return;
      }

      const blob = await res.blob();
      const disposition = res.headers.get("Content-Disposition");
      const match = disposition?.match(/filename="?([^"]+)"?/);
      const filename = match?.[1] ?? "main";

      const url = URL.createObjectURL(blob);
      const anchor = document.createElement("a");
      anchor.href = url;
      anchor.download = filename;
      anchor.click();
      URL.revokeObjectURL(url);

      setOutput(`✓ Binário "${filename}" baixado com sucesso.`);
    } catch (err) {
      setIsError(true);
      setOutput(`Erro ao baixar o binário:\n${err instanceof Error ? err.message : String(err)}`);
    } finally {
      setRunning(false);
    }
  }

  async function handleRun() {
    setRunning(true);
    setIsError(false);
    try {
      const res = await fetch("/api/run", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ code }),
      });

      const data = await res.json() as {
        success: boolean;
        output: string;
        exitCode: number;
        error?: string;
      };

      if (!data.success) {
        setIsError(true);
        setOutput(data.error || data.output || "Erro na execução.");
        return;
      }

      setOutput(data.output);
    } catch (err) {
      setIsError(true);
      setOutput(`Erro ao conectar com o servidor:\n${err instanceof Error ? err.message : String(err)}`);
    } finally {
      setRunning(false);
    }
  }

  function copyOut() {
    navigator.clipboard.writeText(output);
    setCopied(true);
    setTimeout(() => setCopied(false), 1200);
  }

  return (
    <div className="min-h-screen flex flex-col">
      {/* Top bar */}
      <header className="sticky top-0 z-20 backdrop-blur-xl bg-background/70 border-b border-border">
        <div className="mx-auto max-w-[1600px] px-4 sm:px-6 h-16 grid grid-cols-[minmax(0,1fr)_auto] items-center gap-4">
          <div className="flex min-w-0 items-center gap-3">
            <div className="grid h-10 w-10 shrink-0 place-items-center rounded-xl bg-gradient-to-br from-primary to-accent text-primary-foreground shadow-lg shadow-primary/30">
              <Sparkles className="h-5 w-5" />
            </div>
            <div className="min-w-0">
              <h1 className="truncate text-base sm:text-lg font-semibold tracking-tight">
                Minipar Studio
              </h1>
              <p className="truncate text-xs text-muted-foreground">
                IDE | Compiler | Runner
              </p>
            </div>
          </div>
          <nav className="flex items-center gap-1 sm:gap-2">
            <TabButton
              active={tab === "compile"}
              onClick={() => setTab("compile")}
              icon={<Hammer className="h-4 w-4" />}
            >
              Compilar
            </TabButton>
            <TabButton
              active={tab === "execute"}
              onClick={() => setTab("execute")}
              icon={<Play className="h-4 w-4" />}
            >
              Executar
            </TabButton>
          </nav>
        </div>
      </header>

      {/* Main */}
      <main className="flex-1 mx-auto w-full max-w-[1600px] px-4 sm:px-6 py-4 sm:py-6">
        <div className="grid gap-4 lg:gap-6 lg:grid-cols-[240px_minmax(0,1fr)] xl:grid-cols-[260px_minmax(0,1fr)]">
          {/* Sidebar */}
          <aside className="lg:sticky lg:top-20 lg:self-start">
            <div className="rounded-2xl border border-border bg-panel/60 backdrop-blur p-4">
              <div className="flex items-center gap-2 mb-3">
                <FileCode className="h-4 w-4 text-accent" />
                <h2 className="text-sm font-semibold">Exemplos</h2>
              </div>
              <ul className="space-y-1">
                {Object.keys(EXAMPLES).map((name) => (
                  <li key={name}>
                    <button
                      onClick={() => setCode(EXAMPLES[name])}
                      className="group w-full flex items-center justify-between gap-2 rounded-lg px-3 py-2 text-left text-sm text-muted-foreground hover:bg-secondary hover:text-foreground transition"
                    >
                      <span className="truncate">{name}</span>
                      <ChevronRight className="h-3.5 w-3.5 opacity-0 group-hover:opacity-100 transition" />
                    </button>
                  </li>
                ))}
              </ul>
              <div className="mt-6 pt-4 border-t border-border">
                <button
                  onClick={() => setCode("")}
                  className="w-full inline-flex items-center justify-center gap-2 rounded-lg border border-border px-3 py-2 text-xs text-muted-foreground hover:text-foreground hover:bg-secondary transition"
                >
                  <Trash2 className="h-3.5 w-3.5" /> Limpar editor
                </button>
              </div>
            </div>
          </aside>

          {/* Workspace */}
          <section className="grid gap-4 lg:gap-6 lg:grid-cols-2 min-w-0">
            {/* Editor panel */}
            <div className="rounded-2xl border border-border bg-panel/60 backdrop-blur overflow-hidden flex flex-col min-w-0">
              <div className="flex items-center justify-between gap-3 px-4 py-3 border-b border-border">
                <div className="flex items-center gap-2 min-w-0">
                  <Code2 className="h-4 w-4 text-primary shrink-0" />
                  <span className="text-sm font-medium truncate">main.minipar</span>
                  <span className="text-xs text-muted-foreground hidden sm:inline">
                    · {lineCount} linha{lineCount !== 1 ? "s" : ""}
                  </span>
                </div>
                <div className="flex items-center gap-1">
                  <span className="h-2.5 w-2.5 rounded-full bg-warning/80" />
                  <span className="h-2.5 w-2.5 rounded-full bg-success/80" />
                  <span className="h-2.5 w-2.5 rounded-full bg-primary/80" />
                </div>
              </div>

              <div className="relative flex-1 bg-editor scrollbar-thin overflow-auto">
                <div className="flex min-h-full">
                  <pre
                    aria-hidden
                    className="select-none text-right py-4 px-3 text-xs text-muted-foreground/60 font-mono leading-[1.6] border-r border-border"
                  >
                    {Array.from({ length: lineCount }, (_, i) => i + 1).join("\n")}
                  </pre>
                  <div className="editor-scroll flex-1 min-w-0">
                    <Editor
                      value={code}
                      onValueChange={setCode}
                      highlight={(c) =>
                        Prism.highlight(
                          c,
                          Prism.languages.minipar || Prism.languages.plain,
                          "minipar"
                        )
                      }
                      padding={16}
                      style={{
                        minHeight: "100%",
                        fontFamily: "var(--font-mono)",
                        fontSize: 14,
                        color: "var(--foreground)",
                      }}
                    />
                  </div>
                </div>
              </div>

              {/* Action bar */}
              <div className="border-t border-border p-3 sm:p-4 space-y-3">
                {tab === "compile" ? (
                  <>
                    <div className="flex items-center gap-2 text-xs text-muted-foreground">
                      <Settings2 className="h-3.5 w-3.5" />
                      Etapas intermediárias
                    </div>
                    <div className="flex flex-wrap gap-2">
                      {STAGES.map((s) => (
                        <button
                          key={s.key}
                          onClick={() =>
                            setStages((p) => ({ ...p, [s.key]: !p[s.key] }))
                          }
                          className={`px-3 py-1.5 rounded-full text-xs border transition ${
                            stages[s.key]
                              ? "bg-primary text-primary-foreground border-primary shadow shadow-primary/30"
                              : "bg-secondary text-muted-foreground border-border hover:text-foreground"
                          }`}
                        >
                          <span className="mr-1">{s.icon}</span>
                          {s.label}
                        </button>
                      ))}
                    </div>
                    <div className="flex flex-wrap gap-2 pt-1">
                      <button
                        onClick={handleCompile}
                        disabled={running}
                        className="inline-flex items-center gap-2 rounded-xl bg-gradient-to-r from-primary to-accent px-4 py-2.5 text-sm font-semibold text-primary-foreground shadow-lg shadow-primary/30 hover:opacity-95 disabled:opacity-60 transition"
                      >
                        <Hammer className="h-4 w-4" />
                        {running ? "Compilando…" : "Compilar"}
                      </button>
                      <button
                        onClick={handleDownload}
                        disabled={running}
                        className="inline-flex items-center gap-2 rounded-xl border border-border bg-secondary px-4 py-2.5 text-sm hover:bg-secondary/80 disabled:opacity-60 transition"
                      >
                        <Download className="h-4 w-4" />
                        {running ? "Gerando…" : "Baixar binário"}
                      </button>
                    </div>
                  </>
                ) : (
                  <>
                    <button
                      onClick={handleRun}
                      disabled={running}
                      className="inline-flex items-center gap-2 rounded-xl bg-gradient-to-r from-accent to-primary px-4 py-2.5 text-sm font-semibold text-primary-foreground shadow-lg shadow-accent/30 hover:opacity-95 disabled:opacity-60 transition"
                    >
                      <Play className="h-4 w-4" />
                      {running ? "Executando…" : "Executar"}
                    </button>
                  </>
                )}
              </div>
            </div>

            {/* Output panel */}
            <div className="rounded-2xl border border-border bg-panel/60 backdrop-blur overflow-hidden flex flex-col min-h-[400px] min-w-0">
              <div className="flex items-center justify-between gap-3 px-4 py-3 border-b border-border">
                <div className="flex items-center gap-2 min-w-0">
                  {isError ? (
                    <AlertCircle className="h-4 w-4 text-destructive shrink-0" />
                  ) : (
                    <Terminal className="h-4 w-4 text-accent shrink-0" />
                  )}
                  <span className="text-sm font-medium truncate">
                    {tab === "compile" ? "Saída da compilação" : "Saída de execução"}
                  </span>
                  {running && (
                    <span className="ml-2 inline-flex items-center gap-1 text-xs text-muted-foreground">
                      <span className="h-2 w-2 rounded-full bg-warning animate-pulse" />
                      processando…
                    </span>
                  )}
                </div>
                <button
                  onClick={copyOut}
                  className="inline-flex items-center gap-1.5 rounded-md border border-border px-2 py-1 text-xs text-muted-foreground hover:text-foreground hover:bg-secondary transition"
                >
                  {copied ? (
                    <>
                      <Check className="h-3.5 w-3.5 text-success" /> copiado
                    </>
                  ) : (
                    <>
                      <Copy className="h-3.5 w-3.5" /> copiar
                    </>
                  )}
                </button>
              </div>
              <pre
                className={`flex-1 p-4 text-sm font-mono leading-relaxed whitespace-pre-wrap overflow-auto scrollbar-thin bg-editor ${
                  isError ? "text-destructive/90" : "text-foreground/90"
                }`}
              >
                {output}
              </pre>
            </div>
          </section>
        </div>
      </main>

      <footer className="border-t border-border py-4 text-center text-xs text-muted-foreground">
        Minipar Studio · Projeto de Compiladores 2026.1 UFAL
      </footer>
    </div>
  );
}
