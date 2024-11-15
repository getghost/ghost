import { GlossaryClient } from "./client";

export const metadata = {
  title: "Glossary | Ghost",
  description: "Jumpstart your API development with our pre-built solutions.",
  openGraph: {
    title: "Glossary | Ghost",
    description: "Jumpstart your API development with our pre-built solutions.",
    url: "https://ghost.com/glossary",
    siteName: "ghost.com",
    images: [
      {
        url: "https://ghost.com/images/landing/og.png",
        width: 1200,
        height: 675,
      },
    ],
  },
  twitter: {
    title: "Glossary | Ghost",
    card: "summary_large_image",
  },
  icons: {
    shortcut: "/images/landing/ghost.png",
  },
};

export default function GlossaryPage() {
  return (
    <div>
      <GlossaryClient />
    </div>
  );
}
