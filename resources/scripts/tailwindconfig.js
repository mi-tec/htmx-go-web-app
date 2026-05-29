tailwind.config = {
	theme: {
		extend: {
			colors: {
				/* Core surfaces */
				background: '#f7f9fb',
				surface: '#ffffff',
				surfaceAlt: '#f2f4f6',
				surfaceMuted: '#eceef0',

				/* Text */
				foreground: '#191c1e',
				muted: '#45464d',

				/* Brand */
				primary: '#0f172a',      // Deep Medical Blue
				secondary: '#0ea5e9',   // Soft Teal

				/* Borders */
				border: '#e2e8f0',
				outline: '#c6c6cd',

				/* Semantic */
				success: '#15803d',
				warning: '#ca8a04',
				danger: '#ba1a1a',

				/* Sidebar */
				sidebar: '#0f172a',
				sidebarText: 'rgba(255,255,255,0.7)',
				sidebarActive: '#ffffff',
			},

			fontFamily: {
				sans: ['Inter', 'sans-serif'],
			},

			fontSize: {
				'display-lg': ['32px', {
					lineHeight: '40px',
					letterSpacing: '-0.02em',
					fontWeight: '700',
				}],

				'headline-md': ['24px', {
					lineHeight: '32px',
					letterSpacing: '-0.01em',
					fontWeight: '600',
				}],

				'headline-sm': ['20px', {
					lineHeight: '28px',
					fontWeight: '600',
				}],

				'body-lg': ['16px', {
					lineHeight: '24px',
					fontWeight: '400',
				}],

				'body-md': ['14px', {
					lineHeight: '20px',
					fontWeight: '400',
				}],

				'label-md': ['12px', {
					lineHeight: '16px',
					letterSpacing: '0.05em',
					fontWeight: '600',
				}],

				'label-sm': ['11px', {
					lineHeight: '14px',
					fontWeight: '500',
				}],
			},

			borderRadius: {
				sm: '2px',
				DEFAULT: '4px',
				md: '6px',
				lg: '8px',
				xl: '12px',
			},

			spacing: {
				'unit': '4px',
				'gutter': '16px',
				'container': '24px',
				'sidebar': '260px',
				'card': '20px',
			},

			boxShadow: {
				card: '0 1px 3px rgba(0,0,0,0.05)',
				modal: '0 10px 20px rgba(0,0,0,0.10)',
			},
		},
	},
}
