import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Minipar Studio — Modern Compiler IDE",
  description: "A modern, responsive web IDE to write, compile and run Minipar programs.",
  openGraph: {
    title: "Minipar Studio",
    description: "Write, compile and run Minipar in a modern web editor.",
    type: "website",
  },
  twitter: {
    card: "summary",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
        <link
          href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=JetBrains+Mono:wght@400;500;600&display=swap"
          rel="stylesheet"
        />
      </head>
      <body className="antialiased" style={{ fontFamily: "'Inter', ui-sans-serif, system-ui, sans-serif" }}>
        {children}
      </body>
    </html>
  );
}
