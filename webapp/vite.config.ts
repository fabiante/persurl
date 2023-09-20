import { defineConfig } from 'vite';
import preact from '@preact/preset-vite';
import {BaseURL} from "./src/route";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [preact()],
	base: BaseURL,
});
