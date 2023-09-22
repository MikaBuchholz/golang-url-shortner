import styles from "./User.module.css";
import { Navigate, useParams } from "@solidjs/router";
import { createSignal, onMount } from "solid-js";

export interface TrueUrlModel {
  true_url: string;
  view_count: number;
}

const UrlById = () => {
  const params = useParams();

  onMount(async () => {
    const res = await fetch(`http://localhost:8080/api/v1/url/${params.id}`);
    let data = (await res.json()) as TrueUrlModel;
    window.open(`${data.true_url}`);

    return;
  });

  return <></>;
};

export default UrlById;
