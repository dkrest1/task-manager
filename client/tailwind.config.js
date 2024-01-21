/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#3498db'
      },
    },
    btn: {
      primary: {
        backgroundColor: '#000',
        color: '#fff',
        '&:hover': {
          backgroundColor: '#466d7d'
        }
      }
    }
  },
  plugins: [],
}

