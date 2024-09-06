import { z } from "zod";

export const PropertySchema = z.object({
  name: z.string().min(1, "Property name is required"),
  units: z.string().min(1, "Units are required"),
});

export type PropertyType = z.infer<typeof PropertySchema>;
