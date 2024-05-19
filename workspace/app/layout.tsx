import type { Metadata } from "next";
import { Inter } from "next/font/google";

import { Providers } from "./providers";
import { ColorModeScript } from "@chakra-ui/react";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Outis Dashboard",
  description: "Outis Dashboard",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ColorModeScript initialColorMode={'dark'} />
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
