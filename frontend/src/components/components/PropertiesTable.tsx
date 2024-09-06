import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import axios from "axios";
import { useEffect, useState } from "react";

type Property = {
  name: string;
  units: string[];
};

export default function PropertiesTable() {
  const [properties, setProperties] = useState<Property[]>([]);

  useEffect(() => {
    async function fetchProperties() {
      try {
        const response = await axios.get<Property[]>(
          "http://localhost:8080/properties"
        );
        setProperties(response.data);
        console.log(response.data);
      } catch (error) {
        console.error("Error fetching properties:", error);
      }
    }

    fetchProperties();
  }, []);
  return (
    <div className=" flex justify-center align-middle items-center border-black">
      <Table className="w-[40rem] border border-gray-400 border-collapse">
        <TableCaption>A list of the Properties.</TableCaption>
        <TableHeader>
          <TableRow className="text-xl  border-b border-gray-400">
            <TableHead className="border-r border-gray-400  text-black">
              Name
            </TableHead>
            <TableHead className="border-gray-400 text-black">Units</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {properties.map((property, index) => (
            <TableRow key={index} className="h-[2rem] ">
              <TableCell className="border-r border-gray-400">
                {property.name}
              </TableCell>
              <TableCell className="border-gray-400">
                {property.units.join(", ")}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
