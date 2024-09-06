import { Button } from "../ui/button";

export default function AddProperty() {
  return (
    <div className="w-full flex  justify-between items-center  py-6">
      <div className="flex  flex-col  flex-1 gap-2 pr-10 ">
        <label className="flex flex-col  text-[1.3rem]  font-mulish font-medium leading-6 text-black">
          Property Name
          <input
            type="text"
            name="name"
            placeholder={"Name"}
            className="block bg-[#eeede7] w-full border-0 border-b-2 border-[#E3191C] py-1.5 px-2   focus:ring-0 focus:outline-none text-gray-900 sm:text-[1.1rem] text-sm"
          />
        </label>
        <label className="flex flex-col  text-[1.3rem]  font-mulish font-medium leading-6 text-black">
          Units
          <input
            type="text"
            name="units"
            placeholder={`Units (ex: ["bedroom", "bathroom"])`}
            className="block bg-[#eeede7] w-full border-0 border-b-2 border-[#E3191C] py-1.5 px-2   focus:ring-0 focus:outline-none text-gray-900 sm:text-[1.1rem] text-sm"
          />
        </label>
      </div>
      <Button className="self-end bg-[#E3191C] hover:bg-[#E3191C]/90">
        Add Property
      </Button>
    </div>
  );
}
