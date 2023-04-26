import type PocketBase from 'pocketbase';

export type User = (PocketBase['authStore']['model'] & { username: string }) | null;

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			pb: PocketBase;
			user: User;
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
