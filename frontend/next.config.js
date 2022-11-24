/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  exportPathMap: (defaultPathMap) => defaultPathMap,
  env: {
    NEXT_PUBLIC_API_BASE: '/api'
  },
  trailingSlash: true
}

module.exports = nextConfig
