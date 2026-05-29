---
name: Clinical Precision System
colors:
  surface: '#f7f9fb'
  surface-dim: '#d8dadc'
  surface-bright: '#f7f9fb'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f2f4f6'
  surface-container: '#eceef0'
  surface-container-high: '#e6e8ea'
  surface-container-highest: '#e0e3e5'
  on-surface: '#191c1e'
  on-surface-variant: '#45464d'
  inverse-surface: '#2d3133'
  inverse-on-surface: '#eff1f3'
  outline: '#76777d'
  outline-variant: '#c6c6cd'
  surface-tint: '#565e74'
  primary: '#000000'
  on-primary: '#ffffff'
  primary-container: '#131b2e'
  on-primary-container: '#7c839b'
  inverse-primary: '#bec6e0'
  secondary: '#006591'
  on-secondary: '#ffffff'
  secondary-container: '#39b8fd'
  on-secondary-container: '#004666'
  tertiary: '#000000'
  on-tertiary: '#ffffff'
  tertiary-container: '#271901'
  on-tertiary-container: '#98805d'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#dae2fd'
  primary-fixed-dim: '#bec6e0'
  on-primary-fixed: '#131b2e'
  on-primary-fixed-variant: '#3f465c'
  secondary-fixed: '#c9e6ff'
  secondary-fixed-dim: '#89ceff'
  on-secondary-fixed: '#001e2f'
  on-secondary-fixed-variant: '#004c6e'
  tertiary-fixed: '#fcdeb5'
  tertiary-fixed-dim: '#dec29a'
  on-tertiary-fixed: '#271901'
  on-tertiary-fixed-variant: '#574425'
  background: '#f7f9fb'
  on-background: '#191c1e'
  surface-variant: '#e0e3e5'
typography:
  display-lg:
    fontFamily: Inter
    fontSize: 32px
    fontWeight: '700'
    lineHeight: 40px
    letterSpacing: -0.02em
  headline-md:
    fontFamily: Inter
    fontSize: 24px
    fontWeight: '600'
    lineHeight: 32px
    letterSpacing: -0.01em
  headline-sm:
    fontFamily: Inter
    fontSize: 20px
    fontWeight: '600'
    lineHeight: 28px
  body-lg:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: '400'
    lineHeight: 24px
  body-md:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '400'
    lineHeight: 20px
  label-md:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '600'
    lineHeight: 16px
    letterSpacing: 0.05em
  label-sm:
    fontFamily: Inter
    fontSize: 11px
    fontWeight: '500'
    lineHeight: 14px
rounded:
  sm: 0.125rem
  DEFAULT: 0.25rem
  md: 0.375rem
  lg: 0.5rem
  xl: 0.75rem
  full: 9999px
spacing:
  unit: 4px
  container-margin: 24px
  gutter: 16px
  sidebar-width: 260px
  card-padding: 20px
---

## Brand & Style

The design system is engineered for high-stakes healthcare environments where clarity, speed, and trust are paramount. The aesthetic is rooted in **Modern Minimalism** with a focus on functional density. It prioritizes information architecture over decorative elements, ensuring that practitioners can navigate patient data without cognitive friction.

The UI evokes a sense of clinical sterility—clean, organized, and reliable—while maintaining a modern edge through refined typography and precise spacing. It avoids unnecessary ornamentation, using whitespace as a primary tool to separate concerns and reduce the perceived complexity of dense medical records.

## Colors

The palette is anchored by **Deep Medical Blue** (#0f172a), used for primary navigation and high-level structural elements to establish authority. **Clean White** (#ffffff) serves as the primary surface color to maintain a "clinical" feel.

**Soft Teal** (#0ea5e9) is utilized as a focused accent color for primary actions, progress indicators, and active states. Neutral grays range from Slate-50 (#f8fafc) for background fills to Slate-400 (#94a3b8) for secondary text. Semantic colors for status (Success, Warning, Error) follow standard healthcare conventions but are slightly desaturated to fit the professional tone.

## Typography

The design system utilizes **Inter** for all roles. Inter’s tall x-height and clinical legibility make it ideal for reading patient charts and data-heavy tables. 

- **Headlines:** Use SemiBold (600) or Bold (700) with slight negative letter-spacing to maintain a compact, professional look.
- **Body Text:** Primarily uses 14px (body-md) for data density, with 16px reserved for long-form notes.
- **Labels:** Small caps or uppercase treatments are used for table headers and section labels to create clear visual anchors without increasing font size.

## Layout & Spacing

This design system follows a **Fixed-Fluid Hybrid Grid**. The sidebar remains at a fixed width (260px) to provide a stable navigational anchor, while the main content area utilizes a fluid 12-column grid.

- **Desktop:** 24px outer margins and 16px gutters. Content is housed in modular cards.
- **Tablet:** 16px margins; sidebar collapses into a narrow icon-only rail or hamburger menu.
- **Mobile:** 12px margins; cards stack vertically.

Spacing follows a strict 4px base unit. Consistent padding within cards (20px) ensures that data points have breathing room, preventing the UI from feeling cluttered despite high information density.

## Elevation & Depth

Hierarchy is established using **Tonal Layering** combined with **Ambient Shadows**. 

1. **Floor (Background):** Slate-50 (#f8fafc) creates a soft contrast against white cards.
2. **Level 1 (Cards/Surface):** Pure white with a 1px border (#e2e8f0) and a very subtle, diffused shadow (Y: 1px, Blur: 3px, Opacity: 0.05).
3. **Level 2 (Modals/Popovers):** Elevated with a more pronounced shadow (Y: 10px, Blur: 20px, Opacity: 0.1) to focus user attention.

Avoid heavy shadows or "neomorphic" effects. The goal is to simulate physical paper layers on a clean desk.

## Shapes

The design system uses a **Soft** shape language. Standard UI elements (buttons, inputs, cards) use a 0.25rem (4px) corner radius. This provides a modern touch while maintaining the structured, rectilinear feel expected in a professional medical tool. 

- **Large Components (Cards):** Use `rounded-lg` (8px) to soften the transition between sections.
- **Interactive Elements:** Buttons and form inputs strictly follow the 4px rule to ensure they feel like "precise" tools.

## Components

### Data Tables
Tables are the core of the system. Use a flat header with `label-md` typography. Rows should have a height of 52px with a subtle hover state (#f1f5f9). Action buttons within tables should be icon-only or ghost-style to minimize visual noise.

### Sidebar Navigation
The sidebar uses the Primary color (#0f172a) as its background. Nav items should use semi-transparent white text, shifting to pure white with a Teal (#0ea5e9) left-border indicator on the active state.

### Stat Cards
Dashboard cards use a "Metric-Label-Trend" layout. Large numbers use `headline-md`. Trends use a small chip (Success or Error) to show percentage changes.

### Buttons
- **Primary:** Solid Teal (#0ea5e9) with white text.
- **Secondary:** White fill with Slate-200 border and Deep Blue text.
- **Destructive:** Red outline or ghost style for high-risk actions.

### Form Layouts
Inputs must have persistent labels. Use 12px `label-md` placed above the input field. Errors are indicated by a 2px bottom border on the input field and a small helper text below in red.
