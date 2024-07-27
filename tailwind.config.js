/** @type {import('tailwindcss').Config} */
export default {
  content: ["./ui/index.html", "./ui/**/*.{js,jsx}"],
  theme: {
    fontFamily: {
      inter: "Inter, sans-serif",
    },
    extend: {},
  },
  extend: {},
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["dracula"],
  },
}
