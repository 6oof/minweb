/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./views/**/*.go", "./app/middleware/*.go"],
	theme: {
		extend: {
			colors: {
				primary: '#FFB801',
			},
		},
	},
	plugins: [],
};

