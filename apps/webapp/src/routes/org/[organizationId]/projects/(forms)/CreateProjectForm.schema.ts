import { z } from 'zod';
import { FeedbackType } from '$lib/types';

export const schema = z.object({
	name: z.string().min(1).max(64),
	feedbackType: z.enum([FeedbackType.Text, FeedbackType.CNPS, FeedbackType.CSAT])
});
