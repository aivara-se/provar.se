/**
 * A helper type to add a date property to an object type.
 */
export type WithDate<T extends object> = T & { date: string };

/**
 * A helper type for a series with "value" and "date" fields.
 */
export type TimeSeries<T> = Array<WithDate<{ value: T }>>;
