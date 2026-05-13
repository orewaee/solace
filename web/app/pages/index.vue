<template>
    <div class="page">
        <div class="block">
            <Mesh width="360px" height="360px" :blur="32" />
        </div>
    </div>
</template>

<style>
.page {
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}

.block {
    width: 360px;
    height: 600px;
    border-radius: 32px;
    background-color: #111112;
    position: relative;
    overflow: hidden;
}

/* Mesh прибит к низу блока */
.block :deep(.mesh-bg) {
    position: absolute;
    bottom: 0;
    left: 0;
}
</style>
