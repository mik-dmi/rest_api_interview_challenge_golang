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
import AddProperty from "./AddProperty";
import { Button } from "../ui/button";

type Property = {
  name: string;
  units: string[];
};

export default function PropertiesTable() {
  const [properties, setProperties] = useState<Property[]>([]);
  const [refreshTable, setRefreshTable] = useState<boolean>(false);

  const deleteProperty = async (propertyName: string) => {
    try {
      await axios({
        method: "DELETE",
        url: "http://localhost:8080/properties",
        data: {
          name: propertyName,
        },
      });
      setRefreshTable(() => true);
    } catch (error) {
      console.error("Error deleting property:", error);
    }
  };
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
    if (refreshTable) {
      setRefreshTable(false);
    }
  }, [refreshTable]);
  return (
    <div className=" flex  flex-col justify-center align-middle items-center border-black">
      <AddProperty setRefreshTable={setRefreshTable} />
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
              <TableCell className="border-gray-400">
                <Button
                  className="bg-[#E3191C] p-3 hover:bg-[#E3191C]/90"
                  onClick={() => deleteProperty(property.name)}
                >
                  X
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
