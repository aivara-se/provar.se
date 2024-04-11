import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

model FeedbackSearch {
  pageLimit?: float64;
  pageOffset?: float64;
  begTimestamp?: string;
  endTimestamp?: string;
  feedbackType?: FeedbackType[];
  feedbackTags?: Record<string>;
  feedbackMeta?: Record<string>;
  feedbackUser?: Record<string>;
}

model FeedbackToCreate {
  type: FeedbackType;
  time: string;
  data: Record<string>;
  tags?: Record<string>;
  meta?: Record<string>;
  user?: Record<string>;
}

model FeedbackCreate {
  feedbacks: FeedbackToCreate[];
}

@useAuth(BearerAuth)
interface Feedback {
  @post()
  @route("/organization/{organizationId}/feedbacks")
  search(organizationId: string, @body req: FeedbackSearch): { feedbacks: FeedbackDetails[] };

  @post()
  @route("/organization/{organizationId}/feedbacks/count")
  count(organizationId: string, @body req: FeedbackSearch): { total: float64 };

  @post()
  @route("/organization/{organizationId}/feedback")
  create(organizationId: string, @body req: FeedbackCreate): void;

  @route("/organization/{organizationId}/feedback/{feedbackId}")
  details(organizationId: string, feedbackId: string): FeedbackDetails;

  @delete()
  @route("/organization/{organizationId}/feedback/{feedbackId}")
  delete(organizationId: string, feedbackId: string): void;
}
