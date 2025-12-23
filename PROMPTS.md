# Adoption-OS Prompts Library

This file contains reusable prompts for the Adoption-OS project.
Usage: Copy the text inside the code block or reference the header title in chat.

---

## 👹 The "Roast" (Design Critique)
**Goal:** Get brutal, honest feedback on UI/CSS by bypassing the AI's polite filter.

```text
I need a brutal, honest critique of this layout. Assume the persona of a strict Design Lead.

Specifically, tell me what makes this look "amateur" or "unfinished".
If you had to refactor one CSS class to make it look 20% more professional, what would it be?

[PASTE CODE OR SCREENSHOT HERE]

---

# 🎨 UI/UX & Design Audits

## 📋 The "Audit" (Heuristic Evaluation)
**Goal:** A technical breakdown of specific UI failures rather than a subjective rating.

```text
Analyze this UI screenshot. Act as a Senior UX Engineer.
Don't just rate it; instead, conduct a heuristic evaluation focusing on:

1. Mobile Responsiveness: Are buttons large enough for touch targets? Is there horizontal scrolling?
2. Visual Hierarchy: Is the 'Call to Action' distinct from secondary buttons?
3. Consistency: Do these margins match our 16px grid system?

Output a bulleted list of the top 3 critical issues, then create a Plan Artifact to fix #1 immediately.

[PASTE SCREENSHOT HERE]

---

The Accessibility Audit (WCAG)
Goal: Ensure the site is usable for visually impaired users (WCAG 2.1 AA standards).

Analyze the current UI in [TARGET_FILE]. Act as an Accessibility Auditor.
Check color contrast ratios, font sizes, and touch targets against WCAG 2.1 AA standards.
Identify the top 3 violations that would make this hard for a visually impaired user to read, and propose specific CSS fixes.

---

The "Mobile-First" Stress Test
Goal: Prevent layout breaks on small screens (like iPhone SE) before they happen.

I see this layout on my desktop, but I need to ensure it works on an iPhone SE (small screen).
Simulate how these flexbox columns in [TARGET_FILE] will behave on a 320px width.
If they will squash or overflow, write the media query to switch them to a vertical stack.

---

The "Vibe Check" (Style Consistency)
Goal: Catch "magic numbers" and ensure strict adherence to the design system.

Review this new component in [TARGET_FILE].
Does it strictly follow the spacing and shadow tokens used in the rest of the project?
Point out any "magic numbers" (arbitrary pixels like 13px) and replace them with our standard variables.

---

The "Reactivity" Check
Goal: Stop unnecessary re-renders and misuse of ref vs reactive.

Review the reactive state in [TARGET_FILE].
Are we correctly using ref() vs reactive()?
Are there any expensive computed properties that are re-calculating too often?
Refactor to ensure we aren't triggering unnecessary re-renders.

---

The Pinia Migration
Goal: Enforce clean architecture by moving logic out of components and into the store.

I want to move the state logic from [COMPONENT_FILE] into our Pinia store [STORE_FILE].
Please extract the functions and state, create the new action in the store, and update the component to just call the store action.
Keep it strict composition API.

---

The "Lazy Load" Optimizer
Goal: Improve load times on the Raspberry Pi by splitting code chunks.

We need to improve the initial load time on the Raspberry Pi.
Review our router configuration in [ROUTER_FILE].
Are we lazy-loading the 'Admin' and 'Spotlight' views?
If not, rewrite the imports to use dynamic import() syntax.

---

The "SQL Injection" Hunter
Goal: Critical security check for raw SQL vulnerabilities.

Act as a Security Engineer. Review the handler receiving the query in [TARGET_GO_FILE].
Verify that we are using parameterized queries ($1, $2) and NOT string concatenation.
If you find any raw SQL concatenation, block the code and fix it immediately.

---

The "N+1" Query Detector
Goal: Prevent database thrashing that kills Raspberry Pi CPU performance.

Review the [FUNCTION_NAME] function in [TARGET_FILE].
It looks like we are looping through rows and running a second query inside the loop.
This will kill the Raspberry Pi's CPU.
Refactor this to use a single JOIN or a batch query.

---

The "Error Handling" Standard
Goal: Standardize logs so debugging isn't a nightmare.

Review my error logging in [TARGET_FILE].
Currently, I'm just printing 'Error'.
Update the error handling to include:
1. The function name.
2. The timestamp.
3. The specific input that caused the failure.