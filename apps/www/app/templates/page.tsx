import { TemplatesClient } from "./client";

export const metadata = {
  title: "Templates | Ghost",
  description: "Jumpstart your API development with our pre-built solutions.",
  openGraph: {
    title: "Templates | Ghost",
    description: "Jumpstart your API development with our pre-built solutions.",
    url: "https://ghost.com/templates",
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
    title: "Templates | Ghost",
    card: "summary_large_image",
  },
  icons: {
    shortcut: "/images/landing/ghost.png",
  },
};

export default function TemplatesPage() {
  return (
    <div>
      <TemplatesClient />
    </div>
  );
}
