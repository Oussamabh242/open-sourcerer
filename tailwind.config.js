/** @type {import('tailwindcss').Config} */

module.exports = {
  content: [
    './**/*.html',  
    './**/*.templ', // Scanning all .templ files
    './**/*.go',    // Scanning all Go templates
    './**/*.js',    // Scanning all Go templates

  ],
  theme: {
    extend: {
      extend: {
                    colors: {
                        'github-dark': '#0d1117',
                        'github-dark-secondary': '#161b22',
                        'github-border': '#30363d',
                        'github-text': '#c9d1d9',
                        'github-text-secondary': '#8b949e',
                    }
                }
    },
  },
  plugins: [],
}

