/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: ['./index.html', './src/**/*.{vue,js,ts}'],
  theme: {
    extend: {
      colors: {
        jade: {
          50: '#eafff5',
          100: '#cdfee5',
          200: '#a0fad1',
          300: '#63f2b9',
          400: '#25e29d',
          500: '#00bd7e',
          600: '#00a46e',
          700: '#00835c',
          800: '#00674a',
          900: '#00553e'
        },

        background: 'rgb(var(--background) / <alpha-value>)',
        foreground: 'rgb(var(--foreground) / <alpha-value>)',

        card: 'rgb(var(--card) / <alpha-value>)',
        'card-foreground': 'rgb(var(--card-foreground) / <alpha-value>)',

        popover: 'rgb(var(--popover) / <alpha-value>)',
        'popover-foreground': 'rgb(var(--popover-foreground) / <alpha-value>)',

        primary: 'rgb(var(--primary) / <alpha-value>)',
        'primary-foreground': 'rgb(var(--primary-foreground) / <alpha-value>)',

        secondary: 'rgb(var(--secondary) / <alpha-value>)',
        'secondary-foreground': 'rgb(var(--secondary-foreground) / <alpha-value>)',

        muted: 'rgb(var(--muted) / <alpha-value>)',
        'muted-foreground': 'rgb(var(--muted-foreground) / <alpha-value>)',

        accent: 'rgb(var(--accent) / <alpha-value>)',
        'accent-foreground': 'rgb(var(--accent-foreground) / <alpha-value>)',

        destructive: 'rgb(var(--destructive) / <alpha-value>)',
        'destructive-foreground': 'rgb(var(--destructive-foreground) / <alpha-value>)',

        border: 'rgb(var(--border) / <alpha-value>)',
        input: 'rgb(var(--input) / <alpha-value>)',
        ring: 'rgb(var(--ring) / <alpha-value>)',

        status: {
          online: 'rgb(var(--status-online) / <alpha-value>)',
          offline: 'rgb(var(--status-offline) / <alpha-value>)',
          away: 'rgb(var(--status-away) / <alpha-value>)',
          dnd: 'rgb(var(--status-dnd) / <alpha-value>)'
        }
      },
      screens: {
        'hover-hover': { raw: '(hover: hover)' }
      },
      transitionDuration: {
        400: '400ms'
      }
    }
  },
  plugins: []
}
