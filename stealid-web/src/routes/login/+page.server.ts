import type { Actions } from './$types';
import type { PageServerLoad } from './$types';
import { z } from 'zod';
import { message, superValidate } from 'sveltekit-superforms/server';
import { redirect } from '@sveltejs/kit';
import { ClientResponseError } from 'pocketbase';

const loginSchema = z.object({
	username: z.string().min(1).max(32),
	password: z.string().min(8).max(64)
});

export const load = (async (event) => {
	const form = await superValidate(event, loginSchema);
	return {
		form
	};
}) satisfies PageServerLoad;

export const actions = {
	default: async (event) => {
		const form = await superValidate(event, loginSchema);
		if (!form.valid) {
			return message(form, 'Invalid form');
		}

		try {
			await event.locals.pb
				.collection('users')
				.authWithPassword(form.data.username, form.data.password);
		} catch (e) {
			console.log(e);

			if (e instanceof ClientResponseError) {
				return message(form, e.response.message);
			}

			return message(form, 'Something went wrong');
		}

		throw redirect(303, '/dashboard');
	}
} satisfies Actions;
