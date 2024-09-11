import React, { useState, useEffect } from "react";
import { Typeahead } from "react-bootstrap-typeahead";
import "react-bootstrap-typeahead/css/Typeahead.css";
import "react-bootstrap-typeahead/css/Typeahead.bs5.css";
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";
import { Form } from "react-bootstrap";

interface Station {
  name: string;
  line: string[];
  zone: number;
  id: number[];
}

const StationListPage: React.FC = () => {
  const [stations, setStations] = useState<Station[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const [multiSelections, setMultiSelections] = useState<Station[]>([]);

  useEffect(() => {
    const fetchStations = async () => {
      try {
        const response = await fetch(
          "https://backend-134042448147.europe-west2.run.app/stations"
        );
        if (!response.ok) {
          throw new Error("Failed to fetch stations");
        }
        const data: Station[] = await response.json();
        setStations(data);
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (error: any) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchStations();
  }, []);

  if (loading) return <div>Loading stations...</div>;
  if (error) return <div>Error: {error}</div>;

  const selectedCount = multiSelections.length;
  const size = stations.length;
  return (
    <>
      <Form.Group className="mt-3">
        <Form.Label>
          Select the stations you have been to before in London
        </Form.Label>
        <Typeahead
          id="basic-typeahead-multiple"
          labelKey="name"
          multiple
          onChange={(selected) => setMultiSelections(selected as Station[])}
          options={stations}
          placeholder="Choose several stations..."
          selected={multiSelections}
        />
      </Form.Group>
      <div className="selected-items-container">
        Selected items: {selectedCount} out of {size}
      </div>
    </>
  );
};

export default StationListPage;
