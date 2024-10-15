module.exports = {
  mode: 'jit',
  content: ['./src/**/*.{html,ts}'],
  darkMode: 'class', // or 'media' or false
  theme: {
    hljs: {
      theme: 'night-owl',
    },
  },
  variants: {
    extend: {},
  },
  plugins: [require('tailwind-highlightjs')],
  safelist: [
    {
      pattern: /hljs+/,
    },
  ],
};
