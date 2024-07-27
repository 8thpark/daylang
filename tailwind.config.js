/** @type {import('tailwindcss').Config} */
export default {
  content: ["./ui/index.html", "./ui/**/*.{js,jsx}"],
  theme: {
    fontFamily: {
      nunito: "Nunito Sans, sans-serif",
    },
    extend: {},
  },
  extend: {},
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["dracula"],
  },
}
