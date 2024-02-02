/** @type {import('next').NextConfig} */
// const nextConfig = {}

const dns = require("dns")
dns.setDefaultResultOrder("ipv4first")

const path = require('path');
const nextConfig = {
    reactStrictMode: true,
    webpack(config, options) {
    config.resolve.alias['@'] = path.join(__dirname, 'src')
    return config;
    },
}

module.exports = nextConfig