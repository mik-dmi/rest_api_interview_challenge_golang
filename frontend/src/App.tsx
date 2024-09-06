import "./App.css";
import AddProperty from "./components/components/AddProperty";
import PageHeader from "./components/components/PageHeader";
import PropertiesTable from "./components/components/PropertiesTable";

function App() {
  return (
    <main className=" bg-[#eeede7] flex flex-col max-w-[1200px] mx-auto  justify-center items-center  p-6">
      <PageHeader />
      <div className="w-[40rem] flex justify-center  flex-col items-center">
        <AddProperty />
        <PropertiesTable />
      </div>
    </main>
  );
}

export default App;
