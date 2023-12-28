import * as THREE from "three";
import { OrbitControls } from "three/addons/controls/OrbitControls.js";

// Scene
const scene = new THREE.Scene();
const objs = {
    D: { x: 7, y: 9, z: 67, w: 1, h: 1, d: 2, newz: 0 },
    H: { x: 5, y: 1, z: 86, w: 1, h: 4, d: 1, newz: 0 },
    C: { x: 6, y: 3, z: 96, w: 1, h: 1, d: 3, newz: 0 },
    J: { x: 5, y: 6, z: 165, w: 1, h: 1, d: 4, newz: 0 },
    I: { x: 9, y: 5, z: 184, w: 1, h: 2, d: 1, newz: 20 },
    G: { x: 1, y: 3, z: 227, w: 1, h: 3, d: 1, newz: 20 },
    A: { x: 5, y: 2, z: 265, w: 1, h: 4, d: 1, newz: 20 },
    E: { x: 3, y: 2, z: 274, w: 1, h: 2, d: 1, newz: 20 },
    F: { x: 4, y: 6, z: 287, w: 3, h: 1, d: 1, newz: 20 },
    B: { x: 5, y: 5, z: 345, w: 1, h: 4, d: 1, newz: 20 },
};

const pastelColors = [
    0x7f84e8,
    0x7fe88f,
    0xf7e87f,
    0xf7a97f,
    0xf77fe8,
    0x7f8fe8,
    0x7fe8e8,
    0xe87fe8,
    0xe87f7f,
    0xe8a77f,
    0xe8e87f,
];

const pastelColorsHex = []
for (const color of pastelColors) {
    pastelColorsHex.push(color.toString(16))
}

let z = 1
let i = 0
const csize = 512;
for (const [key, value] of Object.entries(objs)) {
    const ctx = document.createElement("canvas").getContext("2d");
    ctx.canvas.width = csize
    ctx.canvas.height = csize
    ctx.fillStyle = "#" + pastelColorsHex[Math.floor(i % pastelColors.length)];
    ctx.fillRect(0, 0, csize, csize);
    ctx.fillStyle = "black";
    ctx.font = Math.floor(csize/4) + "px Arial";
    ctx.fillText(key, csize/2, csize/2);
    const texture = new THREE.CanvasTexture(ctx.canvas);
    texture.needsUpdate = true;

    const geometry = new THREE.BoxGeometry(value.w, value.h, value.d);
    const material = new THREE.MeshBasicMaterial({
        map: texture,
        transparent: true,
        opacity: 1,

    });

    const cube = new THREE.Mesh(geometry, material);
    // we need to offset origin so that the center of the cube is at the x,y,z
    cube.position.set(value.x + value.w / 2, value.y + value.h / 2, value.newz + value.d / 2);

    scene.add(cube);

    z += value.d;
    i++;
}

scene.background = new THREE.Color('gainsboro');

const axesHelper = new THREE.AxesHelper(50);
// x: red, y: green, z: blue
axesHelper.setColors ( 0xff0000, 0x00ff00, 0x0000ff );

scene.add(axesHelper);


// Set up lights
// const ambientLight = new THREE.AmbientLight(0xffffff, 0.6);
// scene.add(ambientLight);

//
const directionalLight = new THREE.DirectionalLight(0xffffff, 0.6);
directionalLight.position.set(10, 20, 0); // x, y, z
scene.add(directionalLight);

// Perspective camera
const aspect = window.innerWidth / window.innerHeight;
const camera = new THREE.PerspectiveCamera(
    45, // field of view in degrees
    aspect, // aspect ratio
    1, // near plane
    100 // far plane
);


// camera.position.set(20, 0, 0)
// camera.lookAt(20, 0, 0)

// so camera looks at 20, 0, 0 and is 20 units away from it
camera.position.set(6, 30, -5)
camera.lookAt(6, 30, -5)

// Renderer
const renderer = new THREE.WebGLRenderer({ antialias: true });
renderer.setSize(window.innerWidth, window.innerHeight);
renderer.render(scene, camera);
renderer.setAnimationLoop(animationLoop);

// Orbit controls
var controls = new OrbitControls(camera, renderer.domElement);

// Add it to HTML
document.body.appendChild(renderer.domElement);

function animationLoop(t) {
    controls.update();
    directionalLight.position.copy(camera.position);
    renderer.render(scene, camera);
    console.log(camera)
}
