CREATE TABLE IF NOT EXISTS public.discs (
    volume integer NOT NULL,
    id serial,
    PRIMARY KEY ("id")
);

ALTER TABLE public.discs OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.virtual_machines (
    id serial,
    name character varying(50) NOT NULL,
    "cpuCount" integer NOT NULL,
    PRIMARY KEY ("id")
);

ALTER TABLE public.virtual_machines
    ADD CONSTRAINT "cpuCount" CHECK (("cpuCount" > 0)) NOT VALID;

ALTER TABLE public.virtual_machines OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.virtual_machine_discs (
    vm_id integer NOT NULL,
    disk_id integer NOT NULL,
    UNIQUE (vm_id, disk_id)
);

ALTER TABLE ONLY public.virtual_machine_discs
    ADD CONSTRAINT "virtual-machine_discs_disc_id_fkey" FOREIGN KEY (disk_id) REFERENCES public.discs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.virtual_machine_discs
    ADD CONSTRAINT "virtual-machine_discs_vm_id_fkey" FOREIGN KEY (vm_id) REFERENCES public.virtual_machines(id) ON DELETE CASCADE;

ALTER TABLE public.virtual_machine_discs OWNER TO postgres;

INSERT INTO public.discs (volume, id) VALUES (2048, 1);
INSERT INTO public.discs (volume, id) VALUES (2048, 2);
INSERT INTO public.discs (volume, id) VALUES (2048, 3);
INSERT INTO public.discs (volume, id) VALUES (4096, 4);
INSERT INTO public.discs (volume, id) VALUES (4096, 5);
INSERT INTO public.discs (volume, id) VALUES (8192, 6);
INSERT INTO public.discs (volume, id) VALUES (8192, 7);
INSERT INTO public.discs (volume, id) VALUES (8192, 8);
INSERT INTO public.discs (volume, id) VALUES (8192, 9);
INSERT INTO public.discs (volume, id) VALUES (8192, 10);

INSERT INTO public.virtual_machines (id, name, "cpuCount") VALUES (1, 'server-1', 4);
INSERT INTO public.virtual_machines (id, name, "cpuCount") VALUES (2, 'server-2', 4);
INSERT INTO public.virtual_machines (id, name, "cpuCount") VALUES (3, 'server-3', 8);

INSERT INTO public.virtual_machine_discs (vm_id, disk_id) VALUES (1, 3);
INSERT INTO public.virtual_machine_discs (vm_id, disk_id) VALUES (1, 6);
INSERT INTO public.virtual_machine_discs (vm_id, disk_id) VALUES (1, 10);
INSERT INTO public.virtual_machine_discs (vm_id, disk_id) VALUES (2, 1);
INSERT INTO public.virtual_machine_discs (vm_id, disk_id) VALUES (2, 5);
INSERT INTO public.virtual_machine_discs (vm_id, disk_id) VALUES (3, 4);

GRANT ALL ON SCHEMA public TO PUBLIC;
