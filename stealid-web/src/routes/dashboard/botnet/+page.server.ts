import type { PageServerLoad } from './$types';

export interface Bot {
	id: string;
	country_code: string;
	ip: string;
	created: string;
}

export const load = (async ({ locals }) => {
	const bots = await locals.pb.collection('bots').getFullList<Bot>(200, { sort: '-created' });
	return {
		bots: structuredClone(bots)
	};
}) satisfies PageServerLoad;
