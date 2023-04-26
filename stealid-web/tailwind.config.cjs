/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				serif: ['Inter', 'sans-series']
			},
			fontSize: {
				'2xl': '1.618rem',
				'3xl': '2.618rem'
			}
		}
	},
	plugins: []
};
