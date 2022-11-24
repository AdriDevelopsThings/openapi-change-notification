/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  exportPathMap: (defaultPathMap) => defaultPathMap,
  env: {
    NEXT_PUBLIC_API_BASE: 'http://localhost:8080/api'
  }
}

module.exports = nextConfig
