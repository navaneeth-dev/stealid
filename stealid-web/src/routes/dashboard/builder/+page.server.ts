import type { Actions, PageServerLoad } from './$types';

export const actions = {
	default: async (event) => {
		await event.locals.pb.collection('builds').create({ user: event.locals.user?.id });
	}
} satisfies Actions;

export interface Build {
	id: string;
}

export const load = (async ({ locals }) => {
	const builds = await locals.pb.collection('builds').getFullList<Build>(200, { sort: '-created' });
	return {
		builds: structuredClone(builds)
	};
}) satisfies PageServerLoad;
