import type { Actions } from './$types';

export const actions = {
	default: async (event) => {
		await event.locals.pb.collection('builds').create({ user: event.locals.user?.id });
	}
} satisfies Actions;
