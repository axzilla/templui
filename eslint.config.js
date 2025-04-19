const htmlPlugin = require("@html-eslint/eslint-plugin");
const eslintHTMLParser = require("@html-eslint/parser");
const alpinejs = require("eslint-plugin-alpinejs");
const tailwind = require("eslint-plugin-html-tailwind");

module.exports = [
  {
    files: ["out/**/*.html"],
    languageOptions: {
      parser: eslintHTMLParser,
    },
    plugins: {
      "@html-eslint": htmlPlugin,
      alpinejs,
      "html-tailwind": tailwind,
    },
    rules: {
      ...htmlPlugin.configs.recommended.rules,
      ...alpinejs.configs.recommended.rules,
      ...tailwind.configs.recommended.rules,
      "html-tailwind/classname-order": "off",
      "alpinejs/no-unused-xrefs": "off",
      "html-tailwind/no-style-attribute": "off",
      "html-tailwind/no-contradicting-classnames": "warn"

    },
  },
];
