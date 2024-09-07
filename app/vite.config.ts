import { defineConfig } from 'vite'
import viteReact from '@vitejs/plugin-react'
import { TanStackRouterVite } from '@tanstack/router-plugin/vite'
import path from "node:path";

const isTest = process.env.NODE_ENV === 'test'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      // '@': path.resolve(__dirname, './src'),
      '@assets': path.resolve(__dirname, './src/assets'),
      // '@components': path.resolve(__dirname, './src/components'),
    },
  },
  plugins: [
    viteReact(),
    !isTest && TanStackRouterVite(),
  ],

})
