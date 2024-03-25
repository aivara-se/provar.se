import { z } from 'zod';

export const schema = z.object({
	name: z.string().min(1).max(64),
	description: z.string().max(512).optional()
});
