import java.util.ArrayList;

public class DBScan {
    
    private double eps;
    private int minPts;
    private ArrayList<TripRecord> trip_data;
    private ArrayList<Cluster> clusters;

    public DBScan(double eps, int minPts, ArrayList<TripRecord> trip_data){
        this.eps = eps;
        this.minPts = minPts;
        this.trip_data = trip_data;
        this.clusters = new ArrayList<Cluster>();
    }

    public ArrayList<Cluster> get_clusters(){
        run_clustering_algorithm();
        return this.clusters;
    }

    public void run_clustering_algorithm(){
        for(TripRecord node: trip_data){
            if(node.get_cluster_id() == -1){
                ArrayList<TripRecord> neighbours = find_neighbours(node);
                if(neighbours.size() >= minPts){
                    Cluster c = new Cluster(clusters.size());
                    c.add_record(node);
                    c.add_all_records(neighbours);
                    ArrayList<TripRecord> mini_neighbours = new ArrayList<TripRecord>();
                    for(TripRecord neighbour: neighbours){
                        if(neighbour.get_cluster_id() == -1){
                            ArrayList<TripRecord> new_neighbours = find_neighbours(neighbour);
                            if(new_neighbours.size() >= minPts){
                                mini_neighbours.addAll(new_neighbours);
                            }
                        }
                    }
                    c.add_all_records(mini_neighbours);
                    clusters.add(c);
                }
            }
        }
    }

    public ArrayList<TripRecord> find_neighbours(TripRecord node) {
        ArrayList<TripRecord> neighbours = new ArrayList<TripRecord>();
        for(TripRecord elem: trip_data){
            if(!elem.equals(node)){
                double distance = get_distance(node, elem);
                if(distance <= eps){
                    neighbours.add(elem);
                }
            }
        }
        return neighbours;
    }

    public double get_distance(TripRecord node, TripRecord elem) {
        double x1, x2, y1, y2;
        x1 = node.pickup_location().get_latitude();
        y1 = node.pickup_location().get_longitude();
        x2 = elem.pickup_location().get_latitude();
        y2 = elem.pickup_location().get_longitude();
        return Math.sqrt(Math.pow(x2 - x1, 2) + Math.pow(y2 - y1, 2));
    }

}
