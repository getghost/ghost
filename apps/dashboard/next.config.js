/** @type {import('next').NextConfig} */

const securityHeaders = [
  {
    key: "X-Frame-Options",
    value: "SAMEORIGIN",
  },
];
const nextConfig = {
  reactStrictMode: true,
  pageExtensions: ["tsx", "mdx", "ts", "js"],
  productionBrowserSourceMaps: true,
  // we're open-source anyways
  experimental: {
    esmExternals: "loose",
  },
  webpack: (config) => {
    config.cache = Object.freeze({
      type: "memory",
    });
    return config;
  },
  transpilePackages: ["@ghost/db", "@ghost/resend", "@ghost/vercel", "@ghost/error", "@ghost/id"],
  eslint: {
    // Warning: This allows production builds to successfully complete even if
    // your project has ESLint errors.
    ignoreDuringBuilds: true,
  },
  async headers() {
    return [
      {
        source: "/(.*)",
        headers: securityHeaders,
      },
    ];
  },
  rewrites: () => [
    {
      source: "/docs",
      destination: "https://ghost.mintlify.dev/docs",
    },
    {
      source: "/docs/:match*",
      destination: "https://ghost.mintlify.dev/docs/:match*",
    },
    {
      source: "/engineering",
      destination: "https://ghost-engineering.mintlify.dev/engineering",
    },
    {
      source: "/engineering/:match*",
      destination: "https://ghost-engineering.mintlify.dev/engineering/:match*",
    },
  ],
};

module.exports = nextConfig;
