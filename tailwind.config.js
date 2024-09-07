/** @type {import('tailwindcss').Config} */

module.exports = {
  content: [
    './**/*.html',  
    './**/*.templ', // Scanning all .templ files
    './**/*.go',    // Scanning all Go templates
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

