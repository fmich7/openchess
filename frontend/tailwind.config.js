/** @type {import('tailwindcss').Config} */

export default {
  mode: "jit",
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: "#3ecf8e",
        "primary-content": "#030b07",
        "primary-dark": "#2baf74",
        "primary-light": "#67d9a6",
        secondary: "#3e7fcf",
        "secondary-content": "#ffffff",
        "secondary-dark": "#2b66af",
        "secondary-light": "#679ad9",
        background: "#191a1a",
        foreground: "#252726",
        border: "#3e4140",
        copy: "#fbfbfb",
        "copy-light": "#d8dad9",
        "copy-lighter": "#a4a8a6",
        success: "#3ecf3e",
        warning: "#cfcf3e",
        error: "#cf3e3e",
        "success-content": "#030b03",
        "warning-content": "#0b0b03",
        "error-content": "#ffffff",
      },
    },
  },
  plugins: [],
};
