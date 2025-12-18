<script setup lang="ts">
import { ref, watch, nextTick, onBeforeUnmount } from 'vue'

const props = withDefaults(
  defineProps<{
    modelValue?: boolean
    size?: number
    ariaLabel?: string
    width?: string
  }>(),
  {
    modelValue: false,
    size: 28,
    ariaLabel: 'Main navigation',
    width: '80vw',
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

const open = ref(!!props.modelValue)
const panelEl = ref<HTMLElement | null>(null)
const prevActive = ref<HTMLElement | null>(null)
let originalOverflow: string | null = null

watch(
  () => props.modelValue,
  async (v) => {
    open.value = v
    if (v) {
      prevActive.value = document.activeElement as HTMLElement
      originalOverflow ??= document.body.style.overflow
      document.body.style.overflow = 'hidden'
      await nextTick()
      panelEl.value?.focus()
    } else {
      document.body.style.overflow = originalOverflow ?? ''
      originalOverflow = null
      prevActive.value?.focus?.()
    }
  },
  { immediate: true },
)

function toggle() {
  emit('update:modelValue', !open.value)
}
function close() {
  emit('update:modelValue', false)
}
function onKeydown(e: KeyboardEvent) {
  if (e.key === 'Escape') {
    e.preventDefault()
    close()
  }
}

onBeforeUnmount(() => {
  if (originalOverflow !== null) document.body.style.overflow = originalOverflow
})
</script>

<template>
  <button
    class="hx-btn"
    type="button"
    :aria-expanded="open"
    :aria-label="open ? 'Close menu' : 'Open menu'"
    @click="toggle"
  >
    <svg
      :width="size"
      :height="size"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-linecap="round"
      stroke-width="2.2"
      :data-open="open"
      aria-hidden="true"
      role="img"
    >
      <line class="hx-top" x1="3" y1="6" x2="21" y2="6" />
      <line class="hx-mid" x1="3" y1="12" x2="21" y2="12" />
      <line class="hx-bot" x1="3" y1="18" x2="21" y2="18" />
    </svg>
  </button>
  <Teleport to="body">
    <transition name="fade">
      <div v-if="open" class="drawer-backdrop" aria-hidden="true" @click="close" />
    </transition>

    <transition name="slide-in-right">
      <aside
        v-if="open"
        class="drawer"
        :style="{ width: width }"
        role="dialog"
        :aria-label="ariaLabel"
        aria-modal="true"
        tabindex="-1"
        ref="panelEl"
        @keydown="onKeydown"
      >
        <header class="drawer-header">
          <h2 class="drawer-title">Menu</h2>
          <button class="drawer-close" @click="close" aria-label="Close menu">
            <svg
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2.2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <line x1="6" y1="6" x2="18" y2="18" />
              <line x1="6" y1="18" x2="18" y2="6" />
            </svg>
          </button>
        </header>

        <nav class="drawer-content">
          <RouterLink to="/" class="nav-link" @click="close">Home</RouterLink>
          <RouterLink to="/about" class="nav-link" @click="close">About</RouterLink>
          <RouterLink to="/adopt" class="nav-link" @click="close">Adopt</RouterLink>
          <RouterLink to="/volunteer" class="nav-link" @click="close">Volunteer</RouterLink>
        </nav>

        <footer class="drawer-footer">
          <RouterLink to="/donate" class="nav-link donate" @click="close">Donate</RouterLink>
        </footer>
      </aside>
    </transition>
  </Teleport>
</template>

<style scoped>
.hx-btn {
  all: unset;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  border-radius: 12px;
  cursor: pointer;
  color: #fff;
}
.hx-btn:focus-visible {
  outline: 2px solid #2f6bd8;
  outline-offset: 2px;
}
svg {
  transition: transform 0.2s ease;
  color: var(--icon, currentColor);
}
line {
  transform-box: fill-box;
  transform-origin: center;
  transition:
    transform 0.22s ease,
    opacity 0.18s ease;
}
.hx-top {
  transform: translateY(0) rotate(0);
}
.hx-mid {
  opacity: 1;
}
.hx-bot {
  transform: translateY(0) rotate(0);
}
svg[data-open='true'] .hx-top {
  transform: translateY(6px) rotate(45deg);
}
svg[data-open='true'] .hx-mid {
  opacity: 0;
}
svg[data-open='true'] .hx-bot {
  transform: translateY(-6px) rotate(-45deg);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.18s ease;
}
.drawer-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.42);
  z-index: 9998;
}

.slide-in-right-enter-from,
.slide-in-right-leave-to {
  transform: translateX(100%);
}
.slide-in-right-enter-active,
.slide-in-right-leave-active {
  transition: transform 0.22s ease;
}
.drawer {
  position: fixed;
  top: 0;
  right: 0;
  height: 100dvh;
  z-index: 9999;
  background: var(--green);
  color: var(--font-color-light);
  box-shadow: -24px 0 48px rgba(0, 0, 0, 0.16);
  border-top-left-radius: 14px;
  border-bottom-left-radius: 14px;
  display: grid;
  grid-template-rows: auto 1fr auto;
  outline: none;
  @media (min-width: 431px) and (max-width: 768px) {
    max-width: 450px;
  }
  @media (min-width: 769px) and (max-width: 1024px) {
  }
}
.drawer-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 14px 16px;
  border-bottom: 1px solid #ececec;
}
.drawer-title {
  margin: 0;
  font-size: 18px;
}
.drawer-close {
  all: unset;
  cursor: pointer;
  padding: 8px;
  border-radius: 10px;
}
.drawer-close:focus-visible {
  outline: 2px solid #2f6bd8;
  outline-offset: 2px;
}
.drawer-content {
  padding: 12px 16px;
  overflow: auto;
  display: flex;
  flex-direction: column;
  gap: 8px;
  a {
    color: var(--font-color-light);
  }
}
.drawer-footer {
  padding: 12px 16px calc(12px + var(--safe-bottom));
  border-top: 1px solid #ececec;
  display: flex;
  justify-content: center;
  width: 100%;
  box-sizing: border-box;
  margin-bottom: 12px;
  a {
    width: 100%;
  }
}

.nav-link {
  padding: 12px 10px;
  border-radius: 12px;
  text-decoration: none;
  color: #16333a;
  font-weight: 600;
}
.nav-link:hover {
  background: var(--green-hover);
}
.nav-link.donate {
  background: var(--blue);
  color: #fff;
  text-align: center;
}
.btn.primary {
  display: inline-block;
  padding: 12px 16px;
  border-radius: 12px;
  background: #2f6bd8;
  color: #fff;
  text-decoration: none;
  font-weight: 700;
}
</style>
