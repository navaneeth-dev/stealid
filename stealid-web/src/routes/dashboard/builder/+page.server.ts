import type { Actions, PageServerLoad } from './$types';

export const actions = {
	build: async (event) => {
		await event.locals.pb.collection('builds').create({ user: event.locals.user?.id });
	},
	download: async (event) => {
		const data = await event.request.formData();
		const buildID = data.get('buildID') as string;
		const record = await event.locals.pb.collection('builds').getOne(buildID);
		const fileToken = await event.locals.pb.files.getToken();
		console.log(fileToken);
		const url = event.locals.pb.files.getUrl(record, record.implant, { token: fileToken });
		return {
			url
		};
	}
} satisfies Actions;

export interface Build {
	id: string;
	created: string;
	updated: string;
	status: string;
}

export const load = (async ({ locals }) => {
	const builds = await locals.pb.collection('builds').getFullList<Build>(200, { sort: '-created' });
	return {
		builds: structuredClone(builds)
	};
}) satisfies PageServerLoad;
