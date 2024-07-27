module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  ignorePatterns: ["dist", ".eslintrc.cjs", "vite.config.js", "tailwind.config.js"],
  parserOptions: { ecmaVersion: "latest", sourceType: "module" },
  settings: { react: { version: "detect" } },
  extends: [
    "airbnb",
    "plugin:react/recommended",
    "plugin:jsx-a11y/recommended",
    "plugin:prettier/recommended",
    "plugin:tailwindcss/recommended",
  ],
  plugins: ["react", "jsx-a11y", "import", "prettier", "tailwindcss"],
  rules: {
    "prettier/prettier": "error",
    "react/jsx-filename-extension": [1, { extensions: [".js", ".jsx"] }],
    "react/prop-types": "error",
    "tailwindcss/classnames-order": "error",
    "tailwindcss/no-custom-classname": "off",
    "tailwindcss/no-contradicting-classname": "error",
    "no-unused-vars": "error",
  },
}
