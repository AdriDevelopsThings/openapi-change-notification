/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  exportPathMap: (defaultPathMap) => defaultPathMap,
  env: {
    NEXT_PUBLIC_API_BASE: 'http://localhost:8080/api',
    NEXT_PUBLIC_HCAPTCHA_SITE_KEY: '836d1df8-1f5b-44c5-8059-03b04c5a5b9d'
  },
  trailingSlash: true
}

module.exports = nextConfig
