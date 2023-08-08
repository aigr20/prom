module.exports = {
  extends: [
    "eslint:recommended",
    "plugin:react-hooks/recommended",
    "plugin:react/recommended",
    "plugin:import/recommended",
    "plugin:jsx-a11y/recommended",
    "plugin:@typescript-eslint/recommended",
    "eslint-config-prettier",
    "plugin:prettier/recommended",
  ],
  settings: {
    react: {
      version: "detect",
    },
    "import/resolver": {
      node: {
        paths: ["src"],
        extensions: [".js", ".jsx", ".ts", ".tsx"],
      },
    },
  },
  rules: {
    semi: ["warn", "always"],
    "no-empty": ["warn", { allowEmptyCatch: false }],
    "prettier/prettier": ["warn"],
    "react/react-in-jsx-scope": [0, "always"],
  },
};
