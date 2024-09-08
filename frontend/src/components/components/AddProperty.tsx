import { Button } from "../ui/button";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import axios from "axios";
import { PropertySchema, PropertyType } from "@/lib/schema";

export default function AddProperty({
  setRefreshTable,
}: {
  setRefreshTable: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<PropertyType>({
    resolver: zodResolver(PropertySchema),
    mode: "all",
  });

  const onSubmit: SubmitHandler<PropertyType> = async (data) => {
    setError(null);
    let unitsArray;
    try {
      unitsArray = JSON.parse(data.units);
      if (!Array.isArray(unitsArray)) {
        throw new Error("Units should be an array");
      }

      unitsArray = unitsArray.map((unit) => unit.toLowerCase().trim());
      const validUnits = ["kitchen", "bathroom", "bedroom", "living-room"];
      const invalidUnits = unitsArray.filter(
        (unit) => !validUnits.includes(unit)
      );

      if (invalidUnits.length > 0) {
        throw new Error("Invalid units: " + invalidUnits.join(", "));
      }
      setRefreshTable(() => true);
    } catch (error) {
      setError('Invalid format for units. Example: ["bedroom", "bathroom"]');
      return;
    }

    const propertyData = {
      name: data.name,
      units: unitsArray,
    };

    try {
      await axios.post("http://localhost:8080/properties", propertyData);
      reset();
    } catch (error) {
      setError("Failed to add property. Please try again.");
    }
  };

  return (
    <div className="w-full flex justify-between items-center py-6">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col flex-1 gap-2 pr-10"
      >
        <label className="flex flex-col text-[1.3rem] font-mulish font-medium leading-6 text-black">
          Property Name
          <input
            type="text"
            {...register("name")}
            placeholder="Name"
            className="block bg-[#eeede7] w-full border-0 border-b-2 border-[#E3191C] py-1.5 px-2 focus:ring-0 focus:outline-none text-gray-900 sm:text-[1.1rem] text-sm"
          />
          {errors.name && <p className="text-red-500">{errors.name.message}</p>}
        </label>

        <label className="flex flex-col text-[1.3rem] font-mulish font-medium leading-6 text-black">
          Units
          <input
            type="text"
            {...register("units")}
            placeholder={`Units (e.g. ["bedroom", "bathroom"])`}
            className="block bg-[#eeede7] w-full border-0 border-b-2 border-[#E3191C] py-1.5 px-2 focus:ring-0 focus:outline-none text-gray-900 sm:text-[1.1rem] text-sm"
          />
          {errors.units && (
            <p className="text-red-500">{errors.units.message}</p>
          )}
        </label>

        {error && <p className="text-red-500 mt-2">{error}</p>}

        <Button
          type="submit"
          className="self-end bg-[#E3191C] hover:bg-[#E3191C]/90"
        >
          Add Property
        </Button>
      </form>
    </div>
  );
}
