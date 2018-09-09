# Loads up the runm-resource database with records that we then use in our PoC
# scenarios

import argparse
import os
import subprocess
import sys

import sqlalchemy as sa

import inventory_profile
import resource_models

_LOG_FORMAT = "%(level)s %(message)s"
_RESOURCE_SCHEMA_FILE = os.path.join(
    os.path.abspath(os.path.dirname(__file__)), 'resource_schema.sql',
)
_INVENTORY_PROFILES_DIR = os.path.join(
    os.path.abspath(os.path.dirname(__file__)), 'inventory-profiles',
)
_DEFAULT_INVENTORY_PROFILE = '1k-shared-compute'


class RunContext(object):
    def __init__(self, args):
        self.args = args

    def status(self, msg):
        sys.stdout.write(msg + " ... ")
        sys.stdout.flush()

    def status_ok(self, ):
        sys.stdout.write("ok\n")
        sys.stdout.flush()

    def status_fail(self, err):
        sys.stdout.write("FAIL\n")
        sys.stdout.flush()
        sys.stderr.write(" error: %s\n" % err)
        sys.stderr.flush()


def _insert_records(tbl, recs):
    try:
        sess = resource_models.get_session()
        for rec in recs:
            ins = tbl.insert().values(**rec)
            sess.execute(ins)
        sess.commit()
        ctx.status_ok()
    except Exception as err:
        ctx.status_fail(err)


def reset_db(ctx):
    ctx.status("resetting resource PoC database")
    db_user = os.environ.get('RUNM_TEST_RESOURCE_DB_USER', 'root')
    db_pass = os.environ.get('RUNM_TEST_RESOURCE_DB_PASS', '')
    cmd = "mysql -u%s -p%s 2>/dev/null < %s" % (
        db_user, db_pass, _RESOURCE_SCHEMA_FILE,
    )
    subprocess.check_output(cmd, shell=True)
    ctx.status_ok()


def create_resource_classes(ctx):
    ctx.status("creating resource classes")
    tbl = resource_models.get_table('resource_classes')

    recs = [
        dict(
            code="runm.cpu.dedicated",
            description="A logical CPU processor associated with a "
                        "single dedicated host CPU processor"
        ),
        dict(
            code="runm.cpu.shared",
            description="A logical CPU processor that may be executed on "
                        "a host CPU processor along with other shared logical "
                        "CPUs"
        ),
        dict(code='runm.memory', description='Bytes of RAM'),
        dict(code='runm.block_storage', description='Bytes of block storage'),
        dict(code='runm.gpu.virtual', description='virtual GPU context'),
    ]
    _insert_records(tbl, recs)


def create_consumer_types(ctx):
    ctx.status("creating consumer types")
    tbl = resource_models.get_table('consumer_types')

    recs = [
        dict(
            code="runm.machine",
            description="A virtual or baremetal machine",
        ),
        dict(
            code="runm.volume",
            description="A persistent volume",
        ),
    ]
    _insert_records(tbl, recs)


def create_capabilities(ctx):
    ctx.status("creating capabilities")
    tbl = resource_models.get_table('capabilities')

    recs = [
        dict(
            code="hw.cpu.x86.avx",
            description="Intel x86 CPU instruction set extensions for AVX",
        ),
        dict(
            code="hw.cpu.x86.avx2",
            description="Intel x86 CPU instruction set extensions for AVX2",
        ),
        dict(
            code="hw.cpu.x86.vmx",
            description="Intel x86 CPU instruction set extensions for VMX",
        ),
        dict(
            code="storage.disk.hdd",
            description="Block storage is on traditional spinning rust",
        ),
        dict(
            code="storage.disk.ssd",
            description="Block storage is on a solid-state drive",
        ),
    ]
    _insert_records(tbl, recs)


def create_distances(ctx):
    ctx.status("creating distance types")
    dt_tbl = resource_models.get_table('distance_types')
    d_tbl = resource_models.get_table('distances')

    recs = [
        dict(
            code="network",
            description="Relative network distances",
            generation=1,
        ),
        dict(
            code="storage",
            description="Relative storage distances",
            generation=1,
        ),
    ]
    _insert_records(dt_tbl, recs)

    ctx.status("creating distances")

    sess = resource_models.get_session()
    sel = sa.select([dt_tbl.c.id]).where(dt_tbl.c.code == "network")
    net_dt_id = sess.execute(sel).fetchone()[0]

    sel = sa.select([dt_tbl.c.id]).where(dt_tbl.c.code == "storage")
    storage_dt_id = sess.execute(sel).fetchone()[0]

    recs = [
        dict(
            type_id=net_dt_id,
            code="local",
            position=0,
            description="Virtually no network latency",
        ),
        dict(
            type_id=net_dt_id,
            code="datacenter",
            position=1,
            description="latency between leaf switch-connected nodes in a DC",
        ),
        dict(
            type_id=net_dt_id,
            code="remote",
            position=2,
            description="WAN latency",
        ),
        dict(
            type_id=storage_dt_id,
            code="local",
            position=0,
            description="Block storage local to host running machine",
        ),
        dict(
            type_id=storage_dt_id,
            code="row",
            position=1,
            description="NAS server shared by row of compute",
        ),
        dict(
            type_id=storage_dt_id,
            code="remote",
            position=2,
            description="External cloud block storage with WAN latency",
        ),
    ]
    _insert_records(d_tbl, recs)


def load_inventory(ctx):
    fp = os.path.join(_INVENTORY_PROFILES_DIR, args.inventory_profile)
    iprof = inventory_profile.InventoryProfile(fp)
    for pg in iprof.iter_provider_groups():
        print pg


def setup_opts(parser):
    inventory_profiles = []
    for fn in os.listdir(_INVENTORY_PROFILES_DIR):
        fp = os.path.join(_INVENTORY_PROFILES_DIR, fn)
        if os.path.isfile(fp) and fn.endswith('.yaml'):
            inventory_profiles.append(fn[0:len(fn) - 5])

    parser.add_argument('--reset', action='store_true',
                        default=True, help="Reset the database entirely.")
    parser.add_argument('--inventory-profile',
                        choices=inventory_profiles,
                        default=_DEFAULT_INVENTORY_PROFILE,
                        help="Inventory profile to use.")


def main(ctx):
    if ctx.args.reset:
        reset_db(ctx)
        create_resource_classes(ctx)
        create_consumer_types(ctx)
        create_capabilities(ctx)
        create_distances(ctx)

    load_inventory(ctx)


if __name__ == '__main__':
    p = argparse.ArgumentParser(description='Load up resource database.')
    setup_opts(p)
    args = p.parse_args()
    ctx = RunContext(args)
    main(ctx)