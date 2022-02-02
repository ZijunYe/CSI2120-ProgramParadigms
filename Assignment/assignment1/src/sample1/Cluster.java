import java.util.ArrayList;

public class Cluster {

    private int cluster_id;
    private GPScoord centroid;
    private ArrayList<TripRecord> cluster;

    public Cluster(int cluster_id){
        this.cluster_id = cluster_id;
        this.centroid = null;
        this.cluster = new ArrayList<TripRecord>();
    }

    public int get_cluster_id(){
        return this.cluster_id;
    }
    
    public GPScoord get_centroid(){
        return this.centroid;
    }

    public void set_centroid(GPScoord centroid){
        this.centroid = centroid;
    }

    public void calculate_centroid(){
        double sumLat = 0;
        double sumLong = 0;
        for(TripRecord tr: cluster){
            sumLat += tr.pickup_location().get_latitude();
            sumLong += tr.pickup_location().get_longitude();
        }
        sumLat /= cluster.size();
        sumLong /= cluster.size();
        set_centroid(new GPScoord(sumLat, sumLong));
    }

    public void add_record(TripRecord node){
        if(!cluster.contains(node)){
            node.set_cluster_id(this.cluster_id);
            cluster.add(node);

        }
    }

    public void add_all_records(ArrayList<TripRecord> nodesList){
        for(TripRecord node: nodesList){
            add_record(node);
        }
    }

    public void print_cluster_data(){
        calculate_centroid();
        System.out.println(
            this.cluster_id + ", " + this.centroid.get_latitude() + " " + 
            this.centroid.get_longitude() + ", " + this.cluster.size()
        );
    }
}
