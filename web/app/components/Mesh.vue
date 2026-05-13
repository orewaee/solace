<template>
    <div ref="rootRef" class="mesh-bg" :style="{ width, height }">
        <canvas ref="maskCanvasRef" class="mesh-bg__mask-canvas" />

        <div ref="videoWrapRef" class="mesh-bg__video-wrap">
            <video
                ref="videoRef"
                class="mesh-bg__video"
                autoplay
                muted
                loop
                playsinline
                preload="auto"
            >
                <source :src="meshVideo" type="video/mp4" />
            </video>
        </div>

        <div class="mesh-bg__blur" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import meshVideo from "~/assets/mesh.mp4";

const props = withDefaults(
    defineProps<{
        width?: string;
        height?: string;
        blur?: number;
    }>(),
    {
        width: "100%",
        height: "200px",
        blur: 20,
    },
);

const rootRef = ref<HTMLDivElement | null>(null);
const videoRef = ref<HTMLVideoElement | null>(null);
const videoWrapRef = ref<HTMLDivElement | null>(null);
const maskCanvasRef = ref<HTMLCanvasElement | null>(null);

let animId = 0;
let t = 0;
let W = 0;
let H = 0;

const WAVE_BASE = 0.8;
const WAVE_AMP_A = 0.06;
const WAVE_AMP_B = 0.025;
const WAVE_FREQ_A = 1.3;
const WAVE_FREQ_B = 2.2;
const WAVE_SPEED = 0.007;

function waveY(xNorm: number, time: number, h: number): number {
    return (
        h * WAVE_BASE +
        Math.sin(xNorm * Math.PI * 2 * WAVE_FREQ_A + time) * h * WAVE_AMP_A +
        Math.sin(xNorm * Math.PI * 2 * WAVE_FREQ_B + time * 1.3 + 0.5) *
            h *
            WAVE_AMP_B
    );
}

const MASK_STEPS = 120;

function drawMask(
    ctx: CanvasRenderingContext2D,
    w: number,
    h: number,
    time: number,
): void {
    ctx.clearRect(0, 0, w, h);
    ctx.beginPath();

    for (let i = 0; i <= MASK_STEPS; i++) {
        const x = (i / MASK_STEPS) * w;
        const y = waveY(i / MASK_STEPS, time, h);
        i === 0 ? ctx.moveTo(x, y) : ctx.lineTo(x, y);
    }

    ctx.lineTo(w, h);
    ctx.lineTo(0, h);
    ctx.closePath();
    ctx.fillStyle = "#fff";
    ctx.fill();
}

function applyCanvasMask(el: HTMLElement, canvas: HTMLCanvasElement): void {
    const url = canvas.toDataURL();
    el.style.webkitMaskImage = `url(${url})`;
    el.style.maskImage = `url(${url})`;
    el.style.webkitMaskSize = "100% 100%";
    el.style.maskSize = "100% 100%";
    el.style.webkitMaskRepeat = "no-repeat";
    el.style.maskRepeat = "no-repeat";
}

function tick(): void {
    t += WAVE_SPEED;

    const canvas = maskCanvasRef.value;
    const wrap = videoWrapRef.value;

    if (canvas && wrap) {
        const ctx = canvas.getContext("2d");
        if (ctx) {
            drawMask(ctx, W, H, t);
            applyCanvasMask(wrap, canvas);
        }
    }

    animId = requestAnimationFrame(tick);
}

onMounted(() => {
    const root = rootRef.value;
    if (!root) return;

    W = root.clientWidth;
    H = root.clientHeight;

    const canvas = maskCanvasRef.value!;
    canvas.width = W;
    canvas.height = H;

    videoRef.value?.play().catch(() => {});
    tick();
});

onUnmounted(() => {
    cancelAnimationFrame(animId);
});
</script>

<style scoped>
.mesh-bg {
    position: absolute;
    bottom: 0;
    left: 0;
    overflow: hidden;
    background: transparent;
}

.mesh-bg__mask-canvas {
    display: none;
}

.mesh-bg__video-wrap {
    position: absolute;
    inset: 0;
    z-index: 0;
}

.mesh-bg__video {
    display: block;
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.mesh-bg__blur {
    position: absolute;
    inset: 0;
    z-index: 1;
    backdrop-filter: blur(v-bind("`${props.blur}px`"));
    -webkit-backdrop-filter: blur(v-bind("`${props.blur}px`"));
}
</style>
