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
} from "lucide-react";

// importação dinâmica do editor para evitar problemas de SSR.
const Editor = dynamic(() => import("react-simple-code-editor"), { ssr: false });

// definir aqui os tokens para a linguagem minipar no Prism, para realce de sintaxe do editor
if (typeof window !== "undefined") {
  Prism.languages.minipar = {
    comment: [/#.*/, /\/\*[\s\S]*?\*\//],
    string: /"(?:\\.|[^"\\])*"/,
    keyword: /\b(?:var|func|return|if|else|while|for|break|continue|true|false)\b/,
    builtin: /\b(?:print|input|number|string|bool|void)\b/,
    function: /\b[a-zA-Z_]\w*(?=\s*\()/,
    number: /\b\d+(?:\.\d+)?\b/,
    operator: /->|==|!=|<=|>=|&&|\|\||[+\-*/%=<>!]/,
    punctuation: /[{}[\];(),.:]/,
  };
}

// aqui ficam os exemplos rapidos de codigos minipar para popular o editor.
const EXAMPLES: Record<string, string> = {
  "Hello World": `# Hello World\nprint("Hello, World!")\nprint("Welcome to Minipar!")\n`,
  "Variáveis & Aritmética": `var x: number = 10\nvar y: number = 5\n\nprint("x + y =", x + y)\nprint("x * y =", x * y)\n`,
  Funções: `func add(a: number, b: number) -> number {\n    return a + b\n}\n\nprint("10 + 5 =", add(10, 5))\n`,
  Loops: `var n: number = 5\nwhile (n >= 0) {\n    print(n)\n    n = n - 1\n}\nprint("Liftoff!")\n`,
  "Fatorial Recursivo": `func fat(n: number) -> number {\n    if (n == 0 || n == 1) { return 1 }\n    else { return n * fat(n - 1) }\n}\n\nprint("5! =", fat(5))\nprint("10! =", fat(10))\n`,
};

type Stage = { key: string; label: string; icon: string };

const STAGES: Stage[] = [
  { key: "tokens", label: "Tokens", icon: "🔤" },
  { key: "ast", label: "AST", icon: "🌳" },
  { key: "semantic", label: "Semântico", icon: "✓" },
  { key: "tac", label: "TAC", icon: "📝" },
  { key: "c", label: "C Code", icon: "⚙️" },
  { key: "asm", label: "Assembly", icon: "🔧" },
];

// ---- ABAIXO SÃO FUNÇÕES MOCKADAS PARA SIMULAR A COMPILAÇÃO E EXECUÇÃO, APENAS PARA DEMONSTRAÇÃO VISUAL ----
// criar api next.js para chamar o cli do minipar via api
function mockStage(stage: string, code: string): string {
  const lines = code.split("\n").filter((l) => l.trim() && !l.trim().startsWith("#"));
  switch (stage) {
    case "tokens":
      return lines
        .slice(0, 6)
        .map(
          (l, i) =>
            `[${String(i + 1).padStart(2, "0")}] ${l
              .trim()
              .split(/\s+/)
              .map((t) => `‹${t}›`)
              .join(" ")}`
        )
        .join("\n");
    case "ast":
      return `Program\n├── Decl(var)\n├── Call(print)\n└── Block`;
    case "semantic":
      return "✓ Tipos consistentes\n✓ Símbolos resolvidos\n✓ Sem warnings";
    case "tac":
      return "t1 = 10\nt2 = 5\nt3 = t1 + t2\nprint t3";
    case "c":
      return `#include <stdio.h>\nint main(){ printf("Hello\\n"); return 0; }`;
    case "asm":
      return ".global _start\n_start:\n  mov r0, #0\n  bx lr";
    default:
      return "";
  }
}

function mockExecute(code: string, stdin: string): string {
  const prints: string[] = [];
  const re = /print\(([^)]*)\)/g;
  let m: RegExpExecArray | null;
  while ((m = re.exec(code))) {
    const args = m[1]
      .split(",")
      .map((a) => a.trim().replace(/^"|"$/g, ""))
      .join(" ");
    prints.push(args);
  }
  const inputs = stdin.split("\n").filter(Boolean);
  if (inputs.length) prints.push(`> stdin: ${inputs.join(" | ")}`);
  return prints.length
    ? prints.join("\n") + "\n\n✓ Programa finalizado (exit 0)"
    : "(sem saída)\n\n✓ Programa finalizado (exit 0)";
}

// ---- fim das funções mockadas ----

// componente de botão para as abas de compilação e execução, com ícones e estilos dinâmicos

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

// função principal do componente Studio, que contém toda a lógica de estado e renderização da interface, incluindo editor de código, opções de compilação, e área de saída.
export default function Studio() {
  const [code, setCode] = useState(EXAMPLES["Hello World"]);
  const [tab, setTab] = useState<"compile" | "execute">("compile");
  const [stages, setStages] = useState<Record<string, boolean>>({
    tokens: false,
    ast: false,
    semantic: false,
    tac: true,
    c: false,
    asm: false,
  });
  const [output, setOutput] = useState<string>(
    "// A saída aparecerá aqui após compilar ou executar."
  );
  const [stdin, setStdin] = useState("");
  const [running, setRunning] = useState(false);
  const [copied, setCopied] = useState(false);

  const lineCount = useMemo(() => code.split("\n").length, [code]);

  // função de chamada do botão compilar. editar quando cli estiver pronto.
  function handleCompile() {
    setRunning(true);
    setTimeout(() => {
      const enabled = STAGES.filter((s) => stages[s.key]);
      const parts: string[] = [];
      parts.push(`▸ Compilação iniciada · ${lineCount} linhas`);
      parts.push(`▸ Etapas: ${enabled.map((s) => s.label).join(", ") || "—"}`);
      parts.push("");
      enabled.forEach((s) => {
        parts.push(`── ${s.icon} ${s.label} ──────────────────────`);
        parts.push(mockStage(s.key, code));
        parts.push("");
      });
      parts.push("✓ Compilado com sucesso.");
      setOutput(parts.join("\n"));
      setRunning(false);
    }, 400);
  }

  // função de chamada do botão executar. editar quando cli estiver pronto.
  function handleRun() {
    setRunning(true);
    setTimeout(() => {
      setOutput(mockExecute(code, stdin));
      setRunning(false);
    }, 350);
  }

  // função para copiar a saída para a área de transferência, com feedback visual de "copiado".
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
                IDE web · compile · execute · explore
              </p>
            </div>
          </div>
          <nav className="flex items-center gap-1 sm:gap-2">
            <TabButton active={tab === "compile"} onClick={() => setTab("compile")} icon={<Hammer className="h-4 w-4" />}>
              Compilar
            </TabButton>
            <TabButton active={tab === "execute"} onClick={() => setTab("execute")} icon={<Play className="h-4 w-4" />}>
              Executar
            </TabButton>
          </nav>
        </div>
      </header>

      {/* Main */}
      <main className="flex-1 mx-auto w-full max-w-[1600px] px-4 sm:px-6 py-4 sm:py-6">
        <div className="grid gap-4 lg:gap-6 lg:grid-cols-[240px_minmax(0,1fr)] xl:grid-cols-[260px_minmax(0,1fr)]">
          {/* Sidebar examples */}
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
            {/* Editor */}
            <div className="rounded-2xl border border-border bg-panel/60 backdrop-blur overflow-hidden flex flex-col min-w-0">
              <div className="flex items-center justify-between gap-3 px-4 py-3 border-b border-border">
                <div className="flex items-center gap-2 min-w-0">
                  <Code2 className="h-4 w-4 text-primary shrink-0" />
                  <span className="text-sm font-medium truncate">main.minipar</span>
                  <span className="text-xs text-muted-foreground hidden sm:inline">
                    · {lineCount} linhas
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
                      Etapas da compilação
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
                        <Hammer className="h-4 w-4" /> Compilar
                      </button>
                      <button className="inline-flex items-center gap-2 rounded-xl border border-border bg-secondary px-4 py-2.5 text-sm hover:bg-muted transition">
                        <Download className="h-4 w-4" /> Baixar .exe
                      </button>
                    </div>
                  </>
                ) : (
                  <>
                    <label className="block text-xs text-muted-foreground">
                      Entrada do programa (stdin)
                    </label>
                    <textarea
                      value={stdin}
                      onChange={(e) => setStdin(e.target.value)}
                      rows={2}
                      placeholder="Um valor por linha…"
                      className="w-full rounded-lg bg-editor border border-border px-3 py-2 text-sm font-mono placeholder:text-muted-foreground/60 focus:outline-none focus:ring-2 focus:ring-ring resize-none"
                    />
                    <button
                      onClick={handleRun}
                      disabled={running}
                      className="inline-flex items-center gap-2 rounded-xl bg-gradient-to-r from-accent to-primary px-4 py-2.5 text-sm font-semibold text-primary-foreground shadow-lg shadow-accent/30 hover:opacity-95 disabled:opacity-60 transition"
                    >
                      <Play className="h-4 w-4" /> Executar
                    </button>
                  </>
                )}
              </div>
            </div>

            {/* Output */}
            <div className="rounded-2xl border border-border bg-panel/60 backdrop-blur overflow-hidden flex flex-col min-h-[400px] min-w-0">
              <div className="flex items-center justify-between gap-3 px-4 py-3 border-b border-border">
                <div className="flex items-center gap-2 min-w-0">
                  <Terminal className="h-4 w-4 text-accent shrink-0" />
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
              <pre className="flex-1 p-4 text-sm font-mono leading-relaxed text-foreground/90 whitespace-pre-wrap overflow-auto scrollbar-thin bg-editor">
                {output}
              </pre>
            </div>
          </section>
        </div>
      </main>

      <footer className="border-t border-border py-6 text-center text-sm text-muted-foreground">
  <p className="font-semibold">Minipar Studio</p>
  <p className="mt-1">IDE e interface de compilação web, para a linguagem de programação Minipar</p>
  <p className="mt-2 text-xs">
    Desenvolvido por: <span className="font-medium">Felipe Lira, Gabriel Seixas, Marcos Mendonça, Wyvian Valença e Ycaro Sales</span>
  </p>
</footer>

    </div>
  );
}
