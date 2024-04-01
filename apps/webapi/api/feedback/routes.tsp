import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

model FeedbackSearch {
  begTimestamp?: string;
  endTimestamp?: string;
  feedbackType?: "text" | "cnps" | "csat";
  feedbackTags?: Record<string>;
  feedbackMeta?: Record<string>;
  feedbackUser?: Record<string>;
}

model FeedbackCreate {
  questionType: string;
  feedbackTime: string;
  feedbackType: "text" | "cnps" | "csat";
  feedbackData: Record<string>;
  feedbackTags: Record<string>;
  feedbackMeta: Record<string>;
  feedbackUser: Record<string>;
}

@useAuth(BearerAuth)
interface Feedback {
  @post()
  @route("/organization/{organizationId}/feedbacks")
  search(organizationId: string, @body req: FeedbackSearch): FeedbackDetails[];

  @post()
  @route("/organization/{organizationId}/feedback")
  create(organizationId: string, @body req: FeedbackCreate): FeedbackDetails;

  @route("/organization/{organizationId}/feedback/{feedbackId}")
  details(organizationId: string, feedbackId: string): FeedbackDetails;

  @delete()
  @route("/organization/{organizationId}/feedback/{feedbackId}")
  delete(organizationId: string, feedbackId: string): void;
}
